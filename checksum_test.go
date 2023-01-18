package checksum

import (
	"encoding/hex"
	"github.com/lxzan/checksum/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	as := assert.New(t)

	as.Equal(8, New().Size())
	as.Equal(64, New().BlockSize())

	var shuffle = func(list []string) []string {
		var n = len(list)
		for k := 0; k < n; k++ {
			var i = internal.AlphabetNumeric.Intn(n)
			var j = internal.AlphabetNumeric.Intn(n)
			list[i], list[j] = list[j], list[i]
		}
		return list
	}

	var arr = []string{
		"02fc5d226d4522bb88ce965540c51b41",
		"046f47e4178064c5a0715895c82932df",
		"0954150028542f32327c",
		"02fc5d226d4522bb88ce965540c51b411",
		"02fc5d226d4522bb88ce965540c51b42",
	}

	t.Run("order", func(t *testing.T) {
		checker1 := New()
		checker1.WriteStrings([]string{arr[0], arr[1], arr[2]})

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

	t.Run("random base64", func(t *testing.T) {
		var list []string
		for i := 0; i < 1000; i++ {
			var n = internal.AlphabetNumeric.Intn(32)
			var s = string(internal.AlphabetNumeric.Generate(n))
			list = append(list, s)
		}

		var h = New()
		for _, s := range list {
			h.Write([]byte(s))
		}
		var code1 = h.SumBase64()

		list = shuffle(list)
		h.Reset()
		for _, s := range list {
			h.WriteString(s)
		}
		var code2 = h.SumBase64()
		as.Equal(code1, code2)
	})

	t.Run("random hex", func(t *testing.T) {
		var list []string
		for i := 0; i < 1000; i++ {
			var n = internal.AlphabetNumeric.Intn(32)
			var s = string(internal.AlphabetNumeric.Generate(n))
			list = append(list, s)
		}

		var h = New()
		for _, s := range list {
			h.Write([]byte(s))
		}
		var code1 = h.SumHex()

		list = shuffle(list)
		h.Reset()
		h.WriteStrings(list)
		var code2 = h.SumHex()
		as.Equal(code1, code2)
	})
}
