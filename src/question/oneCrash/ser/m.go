package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := ":8080"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen err: %s\n", err)
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


	go func() {

		readErrCount := 0
		for {
			data := make([]byte, 10, 10)
			n, err := conn.Read(data)
			if err != nil {
				readErrCount++
				if readErrCount >= 10 {
					return
				}
			}
			fmt.Printf("read: %d element err %s data %s\n", n, err, string(data))
		}
	}()


	go func() {
		writeErrCount := 0
		hadWrite := 0
		for i := 0; i < 100* 10000; i++ {
			data := []byte("hello")
			n, err := conn.Write(data)
			hadWrite += n
			if err != nil {
				writeErrCount++
				if writeErrCount >= 10 {
					return
				}
				fmt.Printf("%d write: %d element err %s data %s\n", hadWrite, n, err, string(data[:]))
			}

		}
	}()

	select {
	}
}
