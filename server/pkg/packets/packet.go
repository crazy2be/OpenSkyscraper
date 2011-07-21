package packets

import (
	"net"
	"os"
	"io"
)

// PlayerConn represents a connection a server maintains with a player. ServerConn, once implemented, will represent the other end of the connection.
type PlayerConn struct {
	conn net.Conn
}

// Does the bare minimum required to read a packet off of the network stream, putting all of the bytes into a raw buffer. In order to do something useful with this packet, you probably want to Parse() it into a key-value structure. This is not done here, because it could be done async (the reading step could not be done async, only one thread should be accessing the network connection at a time). Swallows KEYIDS packets.
func (pc *PlayerConn) NextPacket() (*Packet, os.Error) {
	packet := new(Packet)
	
	buf := make([]byte, 1)
	n, err := pc.conn.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != 1 {
		return nil, os.NewError("Failed to read packet type!")
	}
	packet.Type = buf[0]
	
	buf = make([]byte, 2)
	n, err = pc.conn.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != 2 {
		return nil, os.NewError("Failed to read packet length!")
	}
	packet.length = uint16(buf[0]) << 4 | uint16(buf[1])
	
	packet.RawData = make([]byte, packet.length)
	n, err = io.ReadFull(pc.conn, packet.RawData)
	if err != nil {
		return nil, err
	}
	if packet.Type == KEYIDS {
		err = pc.ParseKeyIDS(packet)
		if err != nil {
			return nil, err
		}
		return pc.NextPacket()
	}
	return packet, nil
}

// STUB! TODO
func (pc *PlayerConn) ParseKeyIDS(packet *Packet) os.Error {
	return nil
}

// Packet represents a single packet sent over a ClientConn or ServerConn, in an unparsed and unfriendly state. Packets can safely be passed off to be proccessed in different goroutines or through channels, as they contain all information necessary to handle them.
type Packet struct {
	Type byte
	length uint16
	RawData []byte
}