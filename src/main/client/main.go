package main

import (
	"flag"
	"log"
	"main/CMD"
	"net"
)

var server *string = flag.String("ser", ":8080", "ser=127.0.0.1:8080")
var concurrence *int = flag.Int("c", 100, "c=60000 (max <= 60000)")

var match = CMD.NewMatch()

func init() {
	match.Register(CMD.C_None, func(message CMD.Message, writeCh chan CMD.Message) {
		writeCh <- CMD.NewUnRegisterMessage(message)
	})
}

func main() {
	// 收到ping报文,立即返回pong报文
	flag.Parse()
	for i := 0; i < *concurrence; i++ {
		go createConn(i)
	}
	select {
	}
}

func createConn(i int) {
	log.Printf("create conn: %d\n", i)
	conn, err := net.Dial("tcp", *server)
	if err != nil {
		log.Printf("create conn err: %s\n", err)
		return
	}

	readCh := make(chan CMD.Message)
	go func() {
		for {
			msg, err := CMD.ReadFrom(conn)
			if err != nil {
				log.Printf("read err: %s\n", err)
				close(readCh)
				return
			}

			readCh <- msg
		}
	}()

	writeCh := make(chan CMD.Message)
	defer close(writeCh)

	go func() {
		for msg := range writeCh {
			err := msg.WriteTo(conn)
			if err != nil {
				log.Printf("write err: %s\n", err)
			}
		}
	}()

	for msg := range readCh {
		match.Process(msg, writeCh)
	}
}