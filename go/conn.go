package xbp

type Conn struct {
	Protocol Protocol
	nextSeq  uint16
}

func NewConnection(protocol Protocol) *Conn {
	return &Conn{
		Protocol: protocol,
	}
}

func (conn *Conn) getNextSeq() uint16 {
	conn.nextSeq++
	return conn.nextSeq
}

func (conn *Conn) sendPacket(flag byte, seq uint16, text string, payload []byte) error {
	return conn.Protocol.SendPacket(&Packet{
		Flag:    flag,
		Seq:     seq,
		Text:    text,
		Payload: payload,
	})
}

func (conn *Conn) SendMessage(text string, payload []byte) error {
	return conn.sendPacket(FlagMessage, conn.getNextSeq(), text, payload)
}

func (conn *Conn) SendRequest(text string, payload []byte) (uint16, error) {
	seq := conn.getNextSeq()
	return seq, conn.sendPacket(FlagRequest, seq, text, payload)
}

func (conn *Conn) SendResponse(seq uint16, text string, payload []byte) error {
	return conn.sendPacket(FlagResponse, seq, text, payload)
}

func (conn *Conn) ReadPacket() (*Packet, error) {
	return conn.Protocol.ReadPacket()
}

func (conn *Conn) Close() error {
	return conn.Protocol.Close()
}
