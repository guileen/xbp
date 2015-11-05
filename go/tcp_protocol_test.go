package xbp

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtocolReal(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:17777")
	assert.Nil(t, err)

	conn1, err1 := net.Dial("tcp", "127.0.0.1:17777")
	assert.Nil(t, err1)
	p1 := NewTcpProtocol(conn1)

	conn2, err2 := listener.Accept()
	assert.Nil(t, err2)
	p2 := NewTcpProtocol(conn2)

	err = p1.SendPacket(&Packet{
		Flag:    FlagResponse,
		Text:    "222",
		Seq:     123,
		Payload: []byte{1, 2, 3, 4, 5, 6},
	})
	assert.Nil(t, err)

	pkt, err := p2.ReadPacket()

	assert.Nil(t, err)
	log.Println("return packet", pkt.Payload)
	assert.Equal(t, FlagResponse, pkt.Flag&FlagResponse)
	assert.Equal(t, "222", pkt.Text)
	assert.EqualValues(t, 123, pkt.Seq)
	assert.EqualValues(t, 6, len(pkt.Payload))
	assert.EqualValues(t, byte(6), pkt.Payload[5])
	err = conn1.Close()
	assert.Nil(t, err)
}
