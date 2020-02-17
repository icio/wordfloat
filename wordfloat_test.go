package wordfloat_test

import (
	"testing"

	"github.com/icio/wordfloat"
)

func TestReadable(t *testing.T) {
	pw, pf := "", float64(0)
	enumerate(make([]byte, 0, 30), 'A', 'z', func(s []byte) {
		curr := wordfloat.Bytes(s, ' ', '~')
		if pf > curr {
			t.Logf("%s\t%f", pw, pf)
			t.Logf("%s\t%f", s, curr)
			t.FailNow()
		}
		pw, pf = string(s), curr
	})
}

func enumerate(buf []byte, min, max byte, ret func([]byte)) {
	l := len(buf)
	down := cap(buf) > l+1
	buf = append(buf, 0)
	for c := min; c <= max; c++ {
		buf = buf[:l+1]
		buf[l] = c
		ret(buf)
		if down {
			enumerate(buf, min, max, ret)
		}
	}
}
