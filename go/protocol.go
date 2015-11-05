package xbp

// TypeBits - bits of sub protocol
// TypeRPC  - type of RPC. Main feature
// TypePing - type of Ping. Keepalive
// TypeHello - type of Hello. Tell the client information related with protocol, like version, zip, supported encoding
const (
	FlagMessage    byte = 0x00
	FlagResponse   byte = 0x80
	FlagRequest    byte = 0x40
	FlagCRC16      byte = 0x20
	FlagSEQ        byte = 0x10
	FlagLenText    byte = 0x0c
	FlagLenPayload byte = 0x03
)

const MaxLength = ^uint32(0)

type Packet struct {
	ClientId int

	Flag byte
	// message sequence
	Seq           uint16
	Size          uint64
	LengthText    uint32
	LengthPayload uint32
	Text          string
	Payload       []byte
}

func (p *Packet) IsRequest() bool {
	return p.Flag&FlagResponse == FlagResponse
}

func (p *Packet) IsResponse() bool {
	return p.Flag&FlagRequest == FlagRequest
}
func (p *Packet) IsMessage() bool {
	return p.Flag&FlagMessage == FlagMessage
}

type Protocol interface {
	ReadPacket() (*Packet, error)
	SendPacket(*Packet) error
	Close() error
}
