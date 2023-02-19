package soupbinTCP

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func generatedServerMsg() []byte {
	var msg []byte
	var msgLength int
	var serverMsg []byte

	//append the type
	typ := uint8('S')
	msg = append(msg, typ)

	//append the payload
	msg = binary.BigEndian.AppendUint32(msg, uint32(4200000009))
	msgLength += len(msg)

	//append the length
	serverMsg = binary.BigEndian.AppendUint16(serverMsg, uint16(msgLength))

	serverMsg = append(serverMsg, msg...)
	return serverMsg
}

var serverMsg = generatedServerMsg()

func TestGetPacket(t *testing.T) {
	b := bytes.NewBuffer(serverMsg)
	r := require.New(t)
	packet, err := GetPacket(b)

	r.NoError(err, "no error")
	r.Equal(serverMsg[2:], packet, "should the same ")
	r.Equal(len(serverMsg[2:]), len(packet))
}
