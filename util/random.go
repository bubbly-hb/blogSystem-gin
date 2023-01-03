package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	data := []byte{}
	for i := 0; i < 26; i++ {
		data = append(data, byte('a'+i))
		data = append(data, byte('A'+i))
	}
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		res[i] = data[rand.Intn(len(data))]
	}
	return string(res)
}
