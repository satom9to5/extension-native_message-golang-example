package main

import (
	"fmt"
	"log"
	"os"
)

type Message struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

var (
	// for debug
	logFile, _ = os.OpenFile("./log/debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger     = log.New(logFile, "", log.LstdFlags|log.Llongfile)
)

func main() {
	recvMes := &Message{}
	var err error

	if err = Receive(recvMes, os.Stdin); err != nil {
		logger.Fatal(err)
	}

	logger.Println(recvMes)

	sendMes := &Message{
		Type:  "response",
		Value: recvMes.Value,
	}

	if err = Send(sendMes, os.Stdout); err != nil {
		logger.Fatal(err)
	}
}

func (m Message) String() string {
	return fmt.Sprintf("Type:%s, Value:%s", m.Type, m.Value)
}
