package main

import "fmt"

type SMS struct {
	senderName  string
	reciverName string
	msg         string
}

func BroadCastToPeople(msgTospouse, msgtome SMS) (totalCost int, err error) {
	spouseMsgcost, err := msgTospouse.sendMsg()
	if err != nil {
		return 0, err
	}
	meMsgCost, err := msgtome.sendMsg()
	if err != nil {
		return 0, err
	}

	return spouseMsgcost + meMsgCost, nil

}

func (sms SMS) sendMsg() (msgCost int, err error) {
	msg := sms.msg
	if len(msg) > 25 {
		return 0, fmt.Errorf("message too long: maximum length is 25 characters")
	}
	return len(msg) * 10, nil
}
