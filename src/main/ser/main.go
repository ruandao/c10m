package main

import (
	"log"
	"main/CMD"
	"main/conn/target"
	"net"
	"sync"
	"time"
)

var match = CMD.NewMatch()
var connManager = sync.Map{}

func init() {
	match.Register(CMD.C_None, func(message CMD.Message, writeCh chan CMD.Message) {
		writeCh <- CMD.NewUnRegisterMessage(message)
	})

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				connManager.Range(func(key, value interface{}) bool {
					t := key.(target.Target)
					t.Write(CMD.NewPingMessage())

					return true
				})
			}
		}
	}()
}

func main() {
	addr := ":8080"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen at %s err %s", addr, err)
	}

	log.Printf("listen at %s\n", addr)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept err: %s\n", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	quit := make(chan struct{})
	defer close(quit)

	t := target.New(conn)
	t.Register(handler)

	connManager.Store(t, nil)
	defer connManager.Delete(t)


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

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				writeCh <- CMD.NewPingMessage()
			case <-quit:
				return
			}
		}
	}()

	// process command
	for msg := range readCh {
		match.Process(msg, writeCh)
	}

}
