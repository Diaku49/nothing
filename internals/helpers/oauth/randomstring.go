package oauth

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomStr() string {
	// seed
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 9)
	for i := range b {
		b[i] = letterBytes[seed.Intn(len(letterBytes))]
	}
	return string(b)
}
