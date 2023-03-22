package itch

//contained implementation of itch-binary de-serialization

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// choose what type to serialize
const ADD_ORDER_MESSAGE = 'A'
const TIMESTAMP_MESSAGE = 'T'

type SerializerMap map[uint8]SerializerFunc
type SerializerFunc func(b []byte, v Message) error

func NewSerializer() SerializerMap {
	fnMap := SerializerMap{}
	fnMap[ADD_ORDER_MESSAGE] = ParseA
	fnMap[TIMESTAMP_MESSAGE] = ParseT

	return fnMap
}

func ParseT(src []byte, v Message) error {
	t, ok := v.(*ItchTime)
	if !ok {
		return errors.New("not timestamp message ")
	}

	length := len(src)
	if length != 5 {
		return fmt.Errorf("wrong size got %v ", length)
	}

	t.Second = binary.BigEndian.Uint32(src)
	return nil
}

// v should pointer to itch.struct
func ParseA(src []byte, v Message) error {
	t, ok := v.(*AddOrder)
	if !ok {
		return fmt.Errorf("not add order message got %T ", v)
	}

	length := len(src)
	if length != 30 {
		return fmt.Errorf("wrong size got %v ", length)
	}

	t.TID = src[0]
	t.Timestamp = binary.BigEndian.Uint32(src[1:5])
	t.OrderNum = binary.BigEndian.Uint64(src[5:13])
	t.OrderVerb = src[13]
	t.Quantity = binary.BigEndian.Uint64(src[14:22])
	t.OrderBook = binary.BigEndian.Uint32(src[22:26])
	t.Price = binary.BigEndian.Uint32(src[26:30])

	return nil
}
