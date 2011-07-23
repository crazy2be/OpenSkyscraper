package packets

import (
	"bytes"
)

// Packet represents a single packet sent over a ClientConn or ServerConn, in an unparsed and unfriendly state. Packets can safely be passed off to be proccessed in different goroutines or through channels, as they contain all information necessary to handle them.
type Packet struct {
	typ byte
	rawData *bytes.Buffer
}

func New(typ byte, rawData *bytes.Buffer) (*Packet) {
	pc := new(Packet)
	pc.typ = typ
	pc.rawData = rawData
	return pc
}

func (p *Packet) Len() int {
	return p.rawData.Len()
}

