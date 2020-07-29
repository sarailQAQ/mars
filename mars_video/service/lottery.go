package service

import (
	"mars/dao"
	"strings"
	"time"
)

type Prize struct {
	PrizeContent string `json:"prize_content"`
	Room string `json:"room"`
	DrawTime *time.Time `json:"draw_time"`
	Number int `json:"number"`
	KewWord string `json:"key_word"`
	Uids map[uint]string
	MessageChan chan *Message
}

type LotteryMnager struct {
	PrizeChan chan *Prize
}

var LManager = LotteryMnager{PrizeChan:make(chan *Prize)}

func (prize *Prize) Receive() {
	for  {
		select {
		case msg := <- prize.MessageChan:
//			fmt.Println("prize",msg)
			if msg.Uid == 0{ return }
			if strings.Compare(msg.Message,prize.KewWord)==0 {
				prize.Uids[msg.Uid] = msg.Usernaem
			}
		}
	}
}

//抽奖时间到时结束协程
func (prize *Prize) Timer() {
//	beginTime := time.Now()
	timer :=time.NewTimer(10*time.Second)//prize.DrawTime.Sub(beginTime)
	<- timer.C
	prize.MessageChan <- &Message{
		Typ:      "",
		Uid:      0,
		Usernaem: "",
		Room:     "",
		Message:  "",
		Color:    0,
	}
	WsManager.EndLottery <- prize
	LManager.PrizeChan <- prize
}

//统一抽奖
func (manager *LotteryMnager) Work() {
	for {
		select {
		case prize := <-manager.PrizeChan :
//			fmt.Println("manager",prize)
			cnt := uint(0)
			var mod uint
			t := len(prize.Uids) / prize.Number
			if t < 1 { t = 1 }
			mod = uint(t)
			for key,val := range prize.Uids {
				if cnt % mod == 0 {
					WsManager.MessageChan <- &Message{
						Typ:      "lottery",
						Uid:      key,
						Usernaem: val,
						Room:     prize.Room,
						Message:  "Congratulations to " + val + " for getting this prize!",
						Color:    0,
					}
					dao.PrzieRecord(key,prize.PrizeContent)  //记录到数据库
				}
				cnt++
			}
		}
	}
}