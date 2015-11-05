package xbp_test

import (
	"log"

	"github.com/guileen/xbp/go"
)

func ExampleListen() {
	listener, err := xbp.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listen")
	conn, _ := listener.Accept()
	pkt, _ := conn.ReadPacket()
	log.Println(pkt)
	// Output:
}

func ExampleDial() {
	conn, err := xbp.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
	}
	conn.SendRequest("hello", []byte("world"))
	// Output:
}
