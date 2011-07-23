package packets

// Packet types
const (
	P_KEEPALIVE = iota
	P_HELLO
	P_TIMEUPDATE
	P_KEYIDS
)

// Type types (values for packets)
const (
	T_BYTE = iota
	T_INT16
	T_INT32
	T_INT64
	T_FLOAT32
	T_FLOAT64
	T_STRING
)