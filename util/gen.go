package util

import (
	"math/rand"
	"time"
)

var char = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

func GenRandNickname() string {
	var result []byte = make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		result = append(result, char[r.Intn(len(char))])
	}
	return string(result)
}
