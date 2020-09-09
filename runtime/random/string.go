package random

import (
	"math/rand"
	"time"
)

const (
	Alpha        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric      = "0123456789"
	AlphaNumeric = Alpha + Numeric
	Symbol       = "~!@#$%^&*()-=_+"
)

func String(length int, dict string) (str string) {
	r := rand.NewSource(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		pos := r.Int63() % int64(len(dict))
		str += string(dict[pos])
	}
	return
}
