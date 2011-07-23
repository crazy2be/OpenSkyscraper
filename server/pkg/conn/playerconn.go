package conn

import (
	"os"
	"io"
	"fmt"
	"net"
	"bytes"
	
	"opensky/packets"
)

// Conn represents a connection between a server and a player. It is a wrapper for net.Conn that provides much more useful methods within the context of the game, and supports reading and writing arbitrary packets.
type Conn struct {
	rawconn net.Conn
	typekeys []Key
}

func New(rawconn net.Conn) (*Conn, os.Error) {
	conn := new(Conn)
	// Make the slice large enough to hold all of the potential key/value pairs.
	conn.typekeys = make([]Key, 0, 65025)
	conn.rawconn = rawconn
	return conn, nil
}

func (conn *Conn) readN(n int) ([]byte, os.Error) {
	buf := make([]byte, n)
	nread, err := conn.rawconn.Read(buf)
	if err != nil {
		return nil, err
	}
	if nread != n {
		return nil, os.NewError(fmt.Sprintf("Failed to read the requested %d bytes, only read %d!", n, nread))
	}
	return buf, nil
}

// Does the bare minimum required to read a packet off of the network stream, putting all of the bytes into a raw buffer. In order to do something useful with this packet, you probably want to Parse() it into a key-value structure. This is not done here, because it could be done async (the reading step could not be done async, only one thread should be accessing the network connection at a time). Swallows KEYIDS packets.
func (conn *Conn) NextPacket() (*packets.Packet, os.Error) {
	typeb, err := conn.readN(1)
	if err != nil {
		return nil, err
	}
	typ := typeb[0]
	
	lengthb, err := conn.readN(2)
	if err != nil {
		return nil, err
	}
	length := uint16(lengthb[0]) << 4 | uint16(lengthb[1])
	
	rawData := bytes.NewBuffer(make([]byte, 0, length))
	n, err := io.Copyn(rawData, conn.rawconn, int64(length))
	if err != nil {
		return nil, err
	}
	if uint16(n) != length {
		return nil, os.NewError(fmt.Sprintf("Packet was of incorrect length! Expected %d, got %d", length, n))
	}
	packet := packets.New(typ, rawData)
	// Swallow KEYIDS packets
	if typ == packets.P_KEYIDS {
		err = conn.parseKeyIDS(packet)
		if err != nil {
			return nil, err
		}
		return conn.NextPacket()
	}
	return packet, nil
}
