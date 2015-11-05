package xbp

type Connection struct {
	Protocol Protocol
	nextSeq  uint16
}

func NewConnection(protocol Protocol) *Connection {
	return &Connection{
		Protocol: protocol,
	}
}

func (conn *Connection) getNextSeq() uint16 {
	conn.nextSeq++
	return conn.nextSeq
}

func (conn *Connection) sendPacket(flag byte, seq uint16, text string, payload []byte) error {
	return conn.Protocol.SendPacket(&Packet{
		Flag:    flag,
		Seq:     seq,
		Text:    text,
		Payload: payload,
	})
}

func (conn *Connection) SendMessage(text string, payload []byte) error {
	return conn.sendPacket(FlagMessage, conn.getNextSeq(), text, payload)
}

func (conn *Connection) SendRequest(text string, payload []byte) error {
	return conn.sendPacket(FlagWaitResponse, conn.getNextSeq(), text, payload)
}

func (conn *Connection) SendResponse(seq uint16, text string, payload []byte) error {
	return conn.sendPacket(FlagResponse, seq, text, payload)
}

func (conn *Connection) ReadPacket() (*Packet, error) {
	return conn.Protocol.ReadPacket()
}

func (conn *Connection) Close() error {
	return conn.Protocol.Close()
}
