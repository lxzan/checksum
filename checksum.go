package checksum

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
)

var byteOrder = binary.LittleEndian

func string2uint64(b string) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func New() *digest {
	d := digest(make([]byte, 8))
	return &d
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

func (c *digest) WriteString(p string) {
	for len(p) >= 8 {
		v := byteOrder.Uint64(*c) ^ string2uint64(p[:8])
		byteOrder.PutUint64(*c, v)
		p = p[8:]
	}

	for i, v := range p {
		(*c)[i] ^= uint8(v)
	}
}

func (c *digest) WriteStrings(list []string) *digest {
	for _, item := range list {
		c.WriteString(item)
	}
	return c
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
	for i, _ := range *c {
		(*c)[i] = 0
	}
}

func (c *digest) Size() int {
	return 8
}

func (c *digest) BlockSize() int {
	return 64
}
