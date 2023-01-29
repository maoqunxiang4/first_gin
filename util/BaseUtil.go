package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	letters := []byte("badjkabcdukbauevyndswncBAKSDBCAIKBCUIEBQNQILUBNRVREL")
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, n)
	for i := range result {
		//len(letters)  确定长度
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
