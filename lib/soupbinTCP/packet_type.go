package soupbinTCP

// soupbinTCP is transport protocol use by stock market exchange
// The SoupBinTCP client and server communicate by exchanging a series of logical
// packets.
// Each SoupBinTCP logical packet has:
// A. a two byte big-endian length that indicates the length of rest of the packet
// (meaning the length of the payload plus the length of the packet type – which is 1)
// B. a single byte header which indicates the packet type
// C. a variable length payload

//represent logical packet the indicates length of the packet and the packet type
type Packet struct {
	Length [2]byte // big-endian
	Type   byte    // packet Type
}

//represent the  heartbeat packet type payload
type HeartbeatPacket struct {
	Packet
}

//represent the message packet  type payload
//client use this packet type to send to the server
type UnsequencedData struct {
	Packet
	Message []byte
}
