package alphavantage

import (
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	keyLen      = 3
)

var (
	n  = 0
	av *Service
)

func randkey(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	if n > 0 {
		time.Sleep(time.Second * 15)
	}
	n++
	return string(b)
}

func init() {
	client := GetClient(randkey(keyLen))

	var err error
	av, err = New(client)
	if err != nil {
		panic(err)
	}
}
