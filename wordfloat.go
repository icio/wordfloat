package wordfloat

func Bytes(word []byte, max, min byte) float64 {
	v := float64(0)
	d := float64(1)
	r := float64(max - min + 2)
	for _, c := range word {
		d /= r
		if c < min || c > max {
			continue
		}
		v += float64(c-min+1) * d
	}
	return v
}
