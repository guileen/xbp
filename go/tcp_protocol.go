package xbp

import (
	"bufio"
	"io"
	"net"
	"reflect"
)

type TcpProtocol struct {
	// Conn
	Conn net.Conn
	// Reader
	Reader *bufio.Reader
	// Writer
	Writer *bufio.Writer
}

func NewTcpProtocol(conn net.Conn) *TcpProtocol {
	if conn == nil || reflect.ValueOf(conn).IsNil() {
		panic("conn should not be nil")
	}
	protocol := newTcpProtocol(conn, conn)
	protocol.Conn = conn
	return protocol
}

func newTcpProtocol(reader io.Reader, writer io.Writer) *TcpProtocol {
	p := &TcpProtocol{
		Reader: bufio.NewReader(reader),
		Writer: bufio.NewWriter(writer),
	}
	return p
}

func (p *TcpProtocol) Close() error {
	if p.Conn != nil || !reflect.ValueOf(p.Conn).IsNil() {
		return p.Conn.Close()
	}
	return nil
}

func (p *TcpProtocol) SendPacket(pk *Packet) error {
	// TODO mutex writer
	if err := WritePacket(p.Writer, pk); err != nil {
		return err
	}
	return nil
}

func (p *TcpProtocol) ReadPacket() (*Packet, error) {
	// TODO mutex reader
	pkt := &Packet{}
	if err := ReadPacket(p.Reader, pkt); err != nil {
		return nil, err
	}
	return pkt, nil
}
