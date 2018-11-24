package target

import (
	"log"
	"main/CMD"
	"net"
)

type Handler interface {
	Process(message CMD.Message, t *Target)
}
type Target struct {
	handler Handler
	quit chan struct{}
	writeCh chan CMD.Message
}

func (t *Target) Write(message CMD.Message) {
	select {
	case <-t.quit:
		return
	case t.writeCh <- message:
	}
}

func (t *Target) Register(handler Handler)  {

}

func New(conn net.Conn, handler Handler) Target {
	readCh := make(chan CMD.Message)

	t := Target{
		handler:handler,
		quit:make(chan struct{}),
		writeCh:make(chan CMD.Message),
	}

	go func() {
		for {
			msg, err := CMD.ReadFrom(conn)
			if err != nil {
				close(t.quit)
				log.Printf("read err: %s", err)
				return
			}

			t.handler.Process(msg)
		}
	}()

	go func() {
		defer conn.Close()

		for msg := range t.writeCh {
			msg.WriteTo(conn)
		}
	}()

	return t
}
