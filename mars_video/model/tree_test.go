package model

import (
	"fmt"
	"testing"
)

func TestAC(t *testing.T) {
	ac := NewAC()
	ac.Add("奖励积分累计安抚快速了解了")
	ac.Add("我比对面弱吗")
	ac.Build()

	fmt.Println(ac.Find("我比对面弱吗"),ac.Find("test"))
	ac.Add("对线也对不过，打团也打不过")
	ac.Add("对面一直进我野区，下路一直叫我去，怎么去啊")
	ac.Add("4396")
	ac.Add("Test")

	ac.Build()
	fmt.Println(ac.Find("Test"),ac.Find("我比对面弱吗"))
}
