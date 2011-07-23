package packets

import (
	"os"
)

// GenericValue represents a value of an unknown type, as obtained from a packet. It can only be converted to the type it was sent as, anything else returns an error. For example, a value sent as an int32 must be retreived as such, attempting to retreive it as a string results in an error.
type genericValue struct {
	typ byte
	rawbytes []byte
}

func (gv *genericValue) GetString() (string, os.Error) {
	if gv.typ != T_STRING {
		return "", os.NewError("Value is not a string type!")
	}
	return "", os.NewError("GetString() not implemented!")
}

func (gv *genericValue) GetByte() (byte, os.Error) {
	return 0, os.NewError("GetByte() not implemented!")
	//...
}