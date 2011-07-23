package conn

import (
	"os"
	"fmt"
	
	"opensky/packets"
)

// A Key represents the information associated with the numeric keys the protocol uses.
type Key struct {
	name string
	typ byte
}

func (pc *Conn) parseKeyIDS(packet *packets.Packet) os.Error {
	for packet.Len() > 0 {
		err := pc.parseKeyID(packet)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pc *Conn) parseKeyID(packet *packets.Packet) (err os.Error) {
	var key Key
	
	// blargh is the key id for the single key id we are reading, which would normally correspond to a name/type pair such as the one we are parsing. Here, however, it should always be equal to the "special" 0xFFFF.
	blargh, err := packet.ReadUint16()
	if err != nil {
		return
	}
	if blargh != 0xFFFF {
		err = os.NewError(fmt.Sprintf("Invalid key ID %d", blargh))
		return
	}
	
	keyName, err := packet.ReadString()
	if err != nil {
		return
	}
	key.name = keyName
	
	keyId, err := packet.ReadUint16()
	if err != nil {
		return
	}
	
	keyType, err := packet.ReadByte()
	if err != nil {
		return
	}
	key.typ = keyType
	
	pc.typekeys[keyId] = key
	return
}