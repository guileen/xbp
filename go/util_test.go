package xbp

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritePacket(t *testing.T) {
	testSize(t, 0)
	testSize(t, 5)
	testSize(t, 255)
	testSize(t, 65535)
}

func testSize(t *testing.T, size int) {
	var b bytes.Buffer
	text := string(bytes.Repeat([]byte("a"), size))
	payload := bytes.Repeat([]byte{1}, size)
	w := bufio.NewWriter(&b)
	err := WritePacket(w, &Packet{
		Text:    text,
		Seq:     1234,
		Payload: payload,
	})
	assert.Nil(t, err)
	log.Println("size:", size)
	for i := 0; i < len(b.Bytes()) && i < 255; i++ {
		fmt.Printf("%02x ", b.Bytes()[i])
	}
	fmt.Println("")

	var pkt Packet
	r := bufio.NewReader(&b)
	err = ReadPacket(r, &pkt)
	assert.Nil(t, err)
	assert.EqualValues(t, len([]byte(text)), pkt.LengthText)
	assert.EqualValues(t, len(payload), pkt.LengthPayload)
	assert.EqualValues(t, 1234, pkt.Seq)
	assert.EqualValues(t, text, pkt.Text)
	assert.EqualValues(t, payload, pkt.Payload)
}
