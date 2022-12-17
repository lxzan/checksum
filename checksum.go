package checksum

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
)

var byteOrder = binary.LittleEndian

func New() digest {
	return make([]byte, 8)
}

type digest []byte

func (c *digest) Write(p []byte) (n int, err error) {
	n = len(p)

	for len(p) >= 8 {
		v := byteOrder.Uint64(*c) ^ byteOrder.Uint64(p[:8])
		byteOrder.PutUint64(*c, v)
		p = p[8:]
	}

	for i, v := range p {
		(*c)[i] ^= v
	}

	return n, nil
}

func (c *digest) Sum(b []byte) []byte {
	return *c
}

func (c *digest) Sum64() uint64 {
	return byteOrder.Uint64(*c)
}

func (c *digest) SumHex() string {
	return hex.EncodeToString(c.Sum(nil))
}

func (c *digest) SumBase64() string {
	return base64.StdEncoding.EncodeToString(c.Sum(nil))
}

func (c *digest) Reset() {
	*c = (*c)[:0]
}

func (c *digest) Size() int {
	return 8
}

func (c *digest) BlockSize() int {
	return 64
}
