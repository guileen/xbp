package xbp_test

import (
	"log"

	"time"

	"github.com/guileen/xbp/go"
)

func ExampleListen() {
	// server
	listener, err := xbp.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Println("listen")
		conn, _ := listener.Accept()
		for i := 0; i < 10000; i++ {
			pkt, _ := conn.ReadPacket()
			log.Println("read", pkt)
		}
	}()

	// client
	conn, err := xbp.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10000; i++ {
		conn.SendRequest("hello", []byte("world"))
	}
	<-time.After(time.Second)
	// Output:
}
