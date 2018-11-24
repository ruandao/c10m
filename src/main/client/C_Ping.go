package main

import (
	"log"
	"main/CMD"
)

func init() {
	match.Register(CMD.C_Ping, func(message CMD.Message, writeCh chan CMD.Message) {
		writeCh <- CMD.NewPongMessage()
		log.Printf("response pong\n")
	})
}
