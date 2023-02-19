package soupbinTCP

import (
	"encoding/binary"
	"io"
	"log"
)

func GetPacket(conn io.Reader) ([]byte, error) {
	/*
	   get the length of the packet from conn [0:2] it is big-endian,
	   create buffer with the length size ,
	   read the rest of the byte [2:] byte into the buffer,
	   return the buffer
	*/

	//each packet has 2 byte big-endian length that indicates the length of rest of the packet
	lengthBuffer := make([]byte, 2)
	_, err := conn.Read(lengthBuffer)
	if err != nil {
		log.Printf("Error reading: %v\n", err)
		return []byte{}, err
	}
	packetLength := binary.BigEndian.Uint16(lengthBuffer)

	//create buffer for hold the payload
	buffer := make([]byte, packetLength)
	conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading: %v\n", err)
		return []byte{}, err
	}
	return buffer, nil
}
