package packets

import (
	"os"
)

func (p *Packet) ReadByte() (byte, os.Error) {
	buf, err := p.readBytes(1)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func (p *Packet) ReadUint16() (uint16, os.Error) {
	buf, err := p.readBytes(2)
	if err != nil {
		return 0, err
	}
	return uint16(buf[0]) << 4 | uint16(buf[1]), nil
}

func (p *Packet) ReadString() (string, os.Error) {
	length, err := p.ReadUint16()
	if err != nil {
		return "", err
	}
	data, err := p.readBytes(int(length))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (p *Packet) readBytes(num int) ([]byte, os.Error) {
	buf := make([]byte, num)
	n, err := p.rawData.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != num {
		return nil, os.NewError("Short read!")
	}
	return buf, nil
}