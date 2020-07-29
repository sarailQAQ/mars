package service

import (
	"mars/dao"
)

type Room struct {
	OwnerId uint
	RoomName,OwnerName string
	Clients map[uint]*Client
	Prize *Prize
	Register, UnRegister chan *Client
	MessageChan chan *Message
	BlackList 	 map[uint]bool
}

func NewRoom(Id uint,RoomName,Name string) *Room {
 	return &Room{
		OwnerId:     Id,
		OwnerName:   Name,
		RoomName:	 RoomName,
		Clients:     make(map[uint]*Client),
		Register:    make(chan *Client),
		UnRegister:  make(chan *Client),
		MessageChan: make(chan *Message),
		BlackList :	 make(map[uint]bool),
	}
}

//这里有个问题，map是用hash实现的，虽然查询效率很高，但是遍历效率很低
//如果有时间的话，可以用BST(RBTree,Treap,Splay等)来实现O(n)遍历，O(log n)插入、查询的map
func (room *Room) Work() {
	room.BlackList = dao.LoadBlackUsers(room.RoomName)

	for {
		select {
		case msg := <- room.MessageChan :
//			fmt.Println("room",msg)
			//黑名单判断
			if room.BlackList[msg.Uid] {
				continue
			}
			//有抽奖活动
			if room.Prize != nil {
				room.Prize.MessageChan <- msg
			}
			for _,c := range room.Clients {
				c.MessageChan <- msg
			}

		case client := <-room.Register :
			if _,ok := room.Clients[client.Uid]; ok {	continue	}
			room.Clients[client.Uid] = client

		case client := <-room.UnRegister:
			if _,ok := room.Clients[client.Uid]; ok {
				close(client.MessageChan)
				delete(room.Clients,client.Uid)
			}
		}
	}
}