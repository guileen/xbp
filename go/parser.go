package xbp

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

func readLengthByPow(reader *bufio.Reader, powOfLength byte) (uint32, error) {
	// read length
	if powOfLength == 0 {
		return 0, nil
	} else if powOfLength == 1 {
		l, err := reader.ReadByte()
		if err != nil {
			return 0, err
		}
		return uint32(l), nil
	} else if powOfLength == 2 {
		var l uint16
		err := binary.Read(reader, binary.BigEndian, &l)
		if err != nil {
			return 0, err
		}
		return uint32(l), nil
	} else if powOfLength == 3 {
		var l uint32
		err := binary.Read(reader, binary.BigEndian, &l)
		if err != nil {
			return 0, err
		}
		return uint32(l), nil
	}
	return 0, errors.New("wrong powOfLength")
}

func readString(reader *bufio.Reader, textLength uint32) (string, error) {
	textBytes := make([]byte, textLength)
	if _, err := io.ReadFull(reader, textBytes); err != nil {
		return "", err
	}
	return string(textBytes), nil
}

func lengthToPow(length uint32) byte {
	if length > 0xffffffff {
		// return 0x04
		panic("length too long")
	} else if length > 0xffff {
		return 0x03
	} else if length > 0xff {
		return 0x02
	} else if length > 0 {
		return 0x01
	} else {
		return 0
	}
}

func writeLength(writer *bufio.Writer, length uint32) error {
	if length == 0 {
		return nil
	}
	var err error
	// write Payload Length
	if length > 0xffff {
		err = binary.Write(writer, binary.BigEndian, uint32(length))
	} else if length > 0xff {
		err = binary.Write(writer, binary.BigEndian, uint16(length))
	} else {
		err = writer.WriteByte(byte(length))
	}
	return err
}

func WritePacket(writer *bufio.Writer, pk *Packet) error {
	if writer == nil {
		return errors.New("No writer")
	}
	if writer.Available() == 0 {
		return errors.New("Writer not available")
	}
	if pk.LengthPayload == 0 {
		pk.LengthPayload = uint32(len(pk.Payload))
	}
	if pk.LengthText == 0 {
		pk.LengthText = uint32(len([]byte(pk.Text)))
	}
	// write Header
	pk.Flag = pk.Flag | lengthToPow(pk.LengthPayload) | (lengthToPow(pk.LengthText) << 2)

	// write Flag
	if err := binary.Write(writer, binary.BigEndian, pk.Flag); err != nil {
		return err
	}

	// write text length
	if err := writeLength(writer, pk.LengthText); err != nil {
		return err
	}

	// write payload length
	if err := writeLength(writer, pk.LengthPayload); err != nil {
		return err
	}

	// TODO calc CRC16
	var crc16 uint16
	// write Seq
	if err := binary.Write(writer, binary.BigEndian, crc16); err != nil {
		return err
	}

	// write Seq
	if err := binary.Write(writer, binary.BigEndian, pk.Seq); err != nil {
		return err
	}

	// write Text
	if _, err := writer.WriteString(pk.Text); err != nil {
		return err
	}

	// flush first
	// MTU = 576
	// IP = 20
	// TCP = 20
	// XBP = 5 - 13
	// 576-20-20-13
	if pk.LengthPayload+pk.LengthText > 523 {
		// flush header first
		if err := writer.Flush(); err != nil {
			return err
		}
	}
	// TODO make Big payload write more effecient
	// write Payload
	if _, err := writer.Write(pk.Payload); err != nil {
		return err
	}
	return writer.Flush()
}

func ReadPacket(reader *bufio.Reader, pkt *Packet) error {
	var err error
	// read Flag
	pkt.Flag, err = reader.ReadByte()
	if err != nil {
		return err
	}
	powOfTextLength := (pkt.Flag & FlagLenText) >> 2
	powOfLength := pkt.Flag & FlagLenPayload

	// read TextLength, BinaryLength
	pkt.LengthText, err = readLengthByPow(reader, powOfTextLength)
	if err != nil {
		return err
	}
	pkt.LengthPayload, err = readLengthByPow(reader, powOfLength)
	if err != nil {
		return err
	}

	// read CRC16
	var crc16 uint16
	err = binary.Read(reader, binary.BigEndian, &crc16)
	if err != nil {
		return err
	}
	// TODO check CRC16

	// read Seq
	var seq uint16
	err = binary.Read(reader, binary.BigEndian, &seq)
	if err != nil {
		return err
	}
	pkt.Seq = seq

	// read Text by LengthText
	if pkt.Text, err = readString(reader, pkt.LengthText); err != nil {
		return err
	}

	if pkt.LengthPayload > 0xffffff {
		return errors.New("buffer too long")
	}

	// read Payload
	pkt.Payload = make([]byte, pkt.LengthPayload)
	if _, err := io.ReadFull(reader, pkt.Payload); err != nil {
		return err
	}
	return nil

}
