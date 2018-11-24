package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("dial err: %s", err)
	}

	data := []byte("ping")
	n, err := conn.Write(data)
	fmt.Printf("write: %d elem err: %s data: %s\n", n, err, string(data))


}
