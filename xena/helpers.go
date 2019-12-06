package xena

import (
	"math/rand"
	"time"
)

const (
	maxIdLen    = 32
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	src = rand.NewSource(time.Now().UnixNano())
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ID(prefix string) string {
	if len(prefix) > maxIdLen {
		return prefix[:maxIdLen]
	}

	n := 32 - len(prefix)
	str := randString(n)
	str = prefix + str
	return str
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
