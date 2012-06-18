package main

import (
	"io"
	"fmt"
	"bytes"
	"encoding/binary"
)

type Packet struct {
	Len uint32
	buf *bytes.Buffer
}

func ReadPacket(rd io.Reader) (p *Packet, err error) {
	p = new(Packet)
	p.buf = new(bytes.Buffer)
	enc := binary.BigEndian
	
	buf := make([]byte, 4)
	_, err = io.ReadFull(rd, buf)
	if err != nil {
		return nil, fmt.Errorf("Could not read packet length: %s", err)
	}
	le := enc.Uint32(buf)
	fmt.Println("Read packet length as", le)
	
	n, err := io.CopyN(p, rd, int64(le))
	if err != nil {
		return nil, fmt.Errorf("Could not read full packet from network: %s. Got %d bytes, expecting %d bytes.", err, n, le)
	}
	
	return p, nil
}

func (p *Packet) Read(buf []byte) (n int, err error) {
	return p.buf.Read(buf)
}

func (p *Packet) Write(buf []byte) (n int, err error) {
	return p.buf.Write(buf)
}

func (p *Packet) Printf(format string, args... interface{}) (int, error) {
	return fmt.Fprintf(p.buf, format, args...)
}