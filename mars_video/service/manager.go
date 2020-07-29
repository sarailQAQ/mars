package service

import (
	"errors"
	"mars/dao"
	"mars/model"
)

type Message struct {
	Typ string `json:"typ"`
	Uid uint `json:"uid"`
	Usernaem string	`json:"usernaem"`
	Room string `json:"room"`
	Message string `json:"message"`
	Color int64 `json:"color"`//字体颜色
}

type Manager struct {
	Filter *model.Filter

	Rooms map[string]*Room
	MessageChan chan *Message
	Register, UnRegister chan *Client
	BeginLottery,EndLottery chan *Prize
	StartRoom chan *Room
}

func (manager *Manager) NewRoom(room *Room) error {
	if _,ok := manager.Rooms[room.RoomName]; ok {
		return errors.New("Room has already exited!")
	}
	manager.Rooms[room.RoomName] = room
	go room.Work()
	return nil
}

var WsManager  = Manager{
	Filter:      model.NewFilter(),
	Rooms:       make(map[string]*Room),
	MessageChan: make(chan *Message),
	Register:	 make(chan *Client),
	UnRegister:  make(chan *Client),
	StartRoom:	 make(chan *Room),
	BeginLottery:make(chan *Prize),
	EndLottery:	 make(chan *Prize),
}

func (manager *Manager) Work() {
	words := dao.LoadSensitveWords()
	for _,v := range words {
		WsManager.Filter.Add(v)
	}
	WsManager.Filter.Build()

	for {
		select {
		case msg := <-manager.MessageChan :
//			fmt.Println("manager",msg)
			if !manager.Filter.Check(msg.Message) {
				continue
			}
			if _,ok := manager.Rooms[msg.Room]; ok {
				manager.Rooms[msg.Room].MessageChan <- msg
			}

		case client := <-manager.Register :
			if _,ok := manager.Rooms[client.Room]; ok {
				manager.Rooms[client.Room].Register <- client
			} else {
				client.MessageChan <- &Message{
					Typ:      "error",
					Uid:      client.Uid,
					Usernaem: client.Username,
					Room:     client.Room,
					Message:  "No such room!",
					Color:    0,
				}
			}

		case client := <-manager.UnRegister:
			if _,ok := manager.Rooms[client.Room]; ok {
				manager.Rooms[client.Room].UnRegister <- client
			}

		case room := <- manager.StartRoom:
			_ = manager.NewRoom(room)

		case prize := <- manager.BeginLottery:
			if _,ok := manager.Rooms[prize.Room]; ok {
				manager.Rooms[prize.Room].Prize = prize
			}

		case prize := <- manager.EndLottery:
			close(prize.MessageChan)
			manager.Rooms[prize.Room].Prize = nil

		}
	}
}
