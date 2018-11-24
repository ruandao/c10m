package main

import "main/CMD"

func init() {
	match.Register(CMD.C_Ping, func(message CMD.Message, writeCh chan CMD.Message) {
		writeCh <- CMD.NewPongMessage()
	})
}
