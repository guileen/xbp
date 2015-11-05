// XBP teXt Binary Protocol.
// Examples:
//  conn, err := Dial("tcp", "12.34.56.78")
//  conn.SendMessage("hello", []byte("world"))
//  conn.SendRequest("hello", []byte("world"))
//  pkt, err := conn.ReadPacket()
//  if pkt.IsRequest() {
//     fmt.Println("Text:%s", pkt.Text)
//     conn.SendResponse(pkt.Sequence, "hello", []byte("world"))
//  }
package xbp

import "net"

// Create new `Connection`
func New(conn net.Conn) *Conn {
	return NewConnection(NewTcpProtocol(conn))
}

// Dial connects to the address on the named network.
// See net.Dial for more information.
func Dial(network, addr string) (*Conn, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
		return New(conn), nil
	}
	return New(conn), nil
}

// Listen announces on the local network address laddr
// The network net must be a stream-oriented network: "tcp", "tcp4",
// "tcp6", "unix" or "unixpacket".
// See net.Dial for the syntax of laddr.
func Listen(network, laddr string) (*Listener, error) {
	netlistener, err := net.Listen(network, laddr)
	if err != nil {
		return nil, err
	}
	return &Listener{
		listener: netlistener,
	}, nil
}

type Listener struct {
	listener net.Listener
}

func (l *Listener) Accept() (*Conn, error) {
	conn, err := l.listener.Accept()
	if err != nil {
		return nil, err
	}
	return New(conn), err
}

func (l *Listener) Close() error {
	return l.listener.Close()
}

func (l *Listener) Addr() net.Addr {
	return l.listener.Addr()
}
