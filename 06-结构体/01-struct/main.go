package main

import (
	"fmt"
	"time"
)

type MessageDetail struct {
	IP   string
	Date time.Time
}

// Message struct
type Message struct {
	sender   string
	body     string
	receiver string
	private  bool
	MessageDetail
}

// type MessageUnread []Message
// type MessageRead []Message

// type MailBox struct {
// 	MessageUnread
// 	MessageRead
// }

func main() {
	m1 := Message{
		sender:   "mike",
		receiver: "alice",
		body:     "hi alice, how are you",
		MessageDetail: MessageDetail{
			IP:   "127.0.0.1",
			Date: time.Now(),
		},
	}
	fmt.Printf("%+v\n", m1)
	// {sender:mike body:hi alice, how are you receiver:alice private:false}

	m2 := Message{"alice", "Fine, Thanks", "mike", false, MessageDetail{"127.0.0.1", time.Now()}}
	fmt.Printf("%+v\n", m2)
	// {sender:alice body:Fine, Thanks receiver:mike private:false}

	m1.private = true
	m2.private = true

	fmt.Printf("m1： ip = %s , date = %s\n", m1.IP, m1.Date)
	// m1： ip = 127.0.0.1 , date = 2019-01-07 03:46:07.243995163 +0800 CST m=+0.000202821

}
