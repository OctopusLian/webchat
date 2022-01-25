package utils

import (
	"bytes"
	"math/rand"
	"time"
)

var letters string = "0123456789abcdefhijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().Unix())
}

func RString(length int) string {
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		buffer.WriteByte(letters[rand.Intn(len(letters))])
	}
	return buffer.String()
}
