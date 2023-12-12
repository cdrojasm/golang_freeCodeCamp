package main

import (
	"fmt"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	senderNumber := mToSend.sender.number
	recipientNumber := mToSend.recipient.number
	senderName := mToSend.sender.name
	recipientName := mToSend.recipient.name
	if senderNumber > 0 && recipientNumber > 0 && senderName != "" && recipientName != "" {
		return true
	}
	return false
}

func test(mToSend messageToSend) {
	canSendThisMessageBool_ := canSendMessage(mToSend)
	var canSendThisMessageStr string
	if canSendThisMessageBool_ {
		canSendThisMessageStr = "Yes"
	} else {
		canSendThisMessageStr = "No"
	}
	fmt.Printf("Sending Message: '%s' to %v\n", mToSend.message, mToSend.recipient.name)
	fmt.Printf("We can send this message ? => %s\n\n", canSendThisMessageStr)
}

func main() {
	usuario_0 := user{
		name:   "miguel",
		number: 976127,
	}
	usuario_1 := user{
		name:   "antonio",
		number: 7980498012734,
	}
	msg_0 := messageToSend{
		message:   "this is a test only",
		sender:    usuario_0,
		recipient: usuario_1,
	}
	msg_1 := messageToSend{
		message: "this is our second test",
		sender:  usuario_0,
	}
	test(msg_0)
	test(msg_1)
}
