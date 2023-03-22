package itch

type Message interface {
	// Map() map[int]interface{}
}

type ItchTime struct {
	Second uint32
}

func (it *ItchTime) Map() map[int]interface{} {
	return nil
}

// []byte{'A',29, 205, 101, 0,0,0,0,0,0,0,0,99,'B',0,0,0,0,0,0,0, 100,0,0,6,236,0,0,78,32}
// []byte{29, 205, 101, 0]		500 000 000 in nanosecond
// []byte{0,0,0,0,0,0,0,99}		99
// []byte{'B'}					'B'
// []byte{0,0,0,0,0,0,0, 100}	100
// []byte{0,0,6,236}			1772
// []byte{0,0,78,32}			20000

type AddOrder struct {
	TID       uint8
	Timestamp uint32 // 4 length, nano-second
	OrderNum  uint64 // 8 length,
	OrderVerb uint8  // 1 length,"B" or "S"
	Quantity  uint64 // 8 length,
	OrderBook uint32 // 4 length,
	Price     uint32 // 4 length,
}

func (ao *AddOrder) Map() map[int]interface{} {
	aoMap := make(map[int]interface{}, 6)
	aoMap[0] = int(ao.Timestamp)
	aoMap[1] = int(ao.OrderNum)
	aoMap[2] = string(ao.OrderVerb)
	aoMap[3] = int(ao.Quantity)
	aoMap[4] = int(ao.OrderBook)
	aoMap[5] = int(ao.Price)
	return aoMap
}
