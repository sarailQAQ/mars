package service

import (
	"github.com/gorilla/websocket"
	"log"

	"mars/model"
)

type Client struct {
	Uid uint
	Username string
	Filter *model.Filter
	Socket *websocket.Conn
	MessageChan chan *Message
	Room string
	Statuse bool
	color int64
}

//读取消息到客户端
func (c *Client) Write() {
	defer func() {
		WsManager.UnRegister <- c
		if err := c.Socket.Close();err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Username, err)
		}
	}()

	for {
		select {
		case msg := <- c.MessageChan:
//			fmt.Println("client",msg)
			err := c.Socket.WriteJSON(msg)
			if err != nil {
				log.Printf("client [%s] writemessage err: %s", c.Username, err)
				return
			}
		}
	}
}

//客户端发送消息到服务端
func (c *Client) Read() {
	defer func() {
		WsManager.UnRegister <- c
		if err := c.Socket.Close();err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Username, err)
		}
	}()

	for  {
		var data  interface{}
		err := c.Socket.ReadJSON(&data)
		if err != nil 	{	break	}
		if dataMap,ok := data.(map[string]interface{}); ok {
			if typ,okk := dataMap["typ"].(string); okk{
				switch typ {
				case "normal":
					if msg,okkk := dataMap["message"].(string); okkk {
						WsManager.MessageChan <- &Message{
							Typ:	  "normal",
							Uid:      c.Uid,
							Usernaem: c.Username,
							Room:     c.Room,
							Message:  msg,
							Color:    c.color,
						}
					}
					
				case "color":
					if col,okkk := dataMap["color"].(int64); okkk {
						c.color = col
					}
				case "Shielding":
					if key,okkk := dataMap["message"].(string); okkk {
						c.Filter.Add(key)
						c.Filter.Build()
					}
				}
			}
			
			
		}
	}
}