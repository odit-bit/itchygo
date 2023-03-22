package itch

// implementation of serializing itch-binary from stream

// type decoder struct {
// 	r          io.Reader
// 	buf        *bufio.Reader
// 	msg        []byte
// 	serializer SerializerMap
// }

// func NewDecoder(r io.Reader) *decoder {
// 	return &decoder{
// 		r:          r,
// 		buf:        bufio.NewReaderSize(r, 128),
// 		msg:        make([]byte, 128),
// 		serializer: NewSerializer(),
// 	}
// }

// func (d *decoder) addOrder(v Message, msg []byte) error {

// 	for i := 0; i < 30; i++ {
// 		b, err := d.buf.ReadByte()
// 		if err != nil {
// 			return err
// 		}
// 		msg[i] = b
// 	}

// 	return d.serializer[ADD_ORDER_MESSAGE](msg, v)
// }

// func (d *decoder) AddOrder(ao *AddOrder) error {
// 	msg := make([]byte, 30)
// 	for i := 0; i < 30; i++ {
// 		b, err := d.buf.ReadByte()
// 		if err != nil {
// 			return err
// 		}
// 		msg[i] = b
// 	}
// 	fn := d.serializer[ADD_ORDER_MESSAGE]
// 	return fn(msg, ao)
// }

// func (d *decoder) timestamp(v Message, msg []byte) error {

// 	for i := 0; i < 5; i++ {
// 		b, err := d.buf.ReadByte()
// 		if err != nil {
// 			return err
// 		}
// 		msg[i] = b
// 	}

// 	return d.serializer[TIMESTAMP_MESSAGE](msg, v)
// }
