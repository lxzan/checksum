package checksum

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	as := assert.New(t)

	var arr = []string{
		"02fc5d226d4522bb88ce965540c51b41",
		"046f47e4178064c5a0715895c82932df",
		"0954150028542f32327c",
		"02fc5d226d4522bb88ce965540c51b411",
		"02fc5d226d4522bb88ce965540c51b42",
	}

	t.Run("order", func(t *testing.T) {
		checker1 := New()
		checker1.Write([]byte(arr[0]))
		checker1.Write([]byte(arr[1]))
		checker1.Write([]byte(arr[2]))

		checker2 := New()
		checker2.Write([]byte(arr[2]))
		checker2.Write([]byte(arr[1]))
		checker2.Write([]byte(arr[0]))

		checker3 := New()
		checker3.Write([]byte(arr[1]))
		checker3.Write([]byte(arr[2]))
		checker3.Write([]byte(arr[0]))

		hash1 := hex.EncodeToString(checker1.Sum(nil))
		hash2 := hex.EncodeToString(checker2.Sum(nil))
		hash3 := hex.EncodeToString(checker3.Sum(nil))
		as.Equal(true, hash1 == hash2)
		as.Equal(true, hash1 == hash3)
	})

	t.Run("value", func(t *testing.T) {
		checker1 := New()
		checker1.Write([]byte(arr[0]))
		checker1.Write([]byte(arr[1]))
		checker1.Write([]byte(arr[2]))

		checker2 := New()
		checker2.Write([]byte(arr[3]))
		checker2.Write([]byte(arr[1]))
		checker2.Write([]byte(arr[2]))

		checker3 := New()
		checker3.Write([]byte(arr[4]))
		checker3.Write([]byte(arr[1]))
		checker3.Write([]byte(arr[2]))

		hash1 := hex.EncodeToString(checker1.Sum(nil))
		hash2 := hex.EncodeToString(checker2.Sum(nil))
		hash3 := hex.EncodeToString(checker3.Sum(nil))
		as.Equal(false, hash1 == hash2)
		as.Equal(false, hash1 == hash3)
	})

	t.Run("sum64", func(t *testing.T) {
		checker1 := New()
		checker1.Write([]byte(arr[0]))
		checker1.Write([]byte(arr[1]))
		checker1.Write([]byte(arr[2]))

		checker2 := New()
		checker2.Write([]byte(arr[2]))
		checker2.Write([]byte(arr[1]))
		checker2.Write([]byte(arr[0]))

		checker3 := New()
		checker3.Write([]byte(arr[4]))
		checker3.Write([]byte(arr[1]))
		checker3.Write([]byte(arr[2]))

		hash1 := checker1.Sum64()
		hash2 := checker2.Sum64()
		hash3 := checker3.Sum64()
		as.Equal(true, hash1 == hash2)
		as.Equal(false, hash1 == hash3)
	})
}
