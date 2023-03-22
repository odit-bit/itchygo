package processor

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/odit-bit/ithcygo/itch"
)

type bidOffer struct {
	serializer itch.SerializerFunc
	// orders     map[int][]itch.AddOrder
	buf []byte
}

func NewBidOffer() *bidOffer {
	bo := &bidOffer{
		serializer: itch.NewSerializer()[itch.ADD_ORDER_MESSAGE],
		// orders:     map[int][]itch.AddOrder{},
		buf: make([]byte, 30),
	}
	return bo
}

func (bo *bidOffer) Read(r io.Reader) (itch.AddOrder, error) {
	buf := bo.buf //make([]byte, 30)
	var ao itch.AddOrder
	n, err := r.Read(buf)
	if n > 0 {
		err = bo.serializer(buf, &ao)
		if err != nil {
			return ao, err
		}

	}
	if err != nil {
		if err == io.EOF {
			return ao, err
		}
		return ao, err
	}
	return ao, nil
}

func (bo *bidOffer) JSONPipeline(r io.Reader, w *bytes.Buffer) error {
	// in := make(chan []byte, 1)
	// errC := make(chan error, 1)
	buf := bo.buf //make([]byte, 30)
	ao := &itch.AddOrder{}

	for {
		n, err := r.Read(buf)
		if n > 0 {
			err = bo.serializer(buf, ao)
			if err != nil {
				return err
			}
			b, err := json.Marshal(ao)
			if err != nil {
				return err
			}
			n, err = w.Write(b)
			log.Println("write n byte", n, "total byte", w.Len())
			if err != nil {
				return err
			}
		}
		if err != nil {
			if err == io.EOF {
				return err
			}
			return err
		}
	}

}
