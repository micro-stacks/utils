package strings

import (
	"math/rand"
	"time"
)

const (
	numbers    = "0123456789"
	numLetters = numbers + "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandNum 生成长度为 len 的数字字符串。
func RandNum(len int) string {
	res := make([]byte, len)
	asset, remain := rand.Uint64(), 16

	for i := 0; i < len; i++ {
		if remain == 0 {
			asset, remain = rand.Uint64(), 16
		}
		if bits4 := asset & 0b1111; bits4 < 10 {
			res[i] = numbers[bits4]
		} else {
			i--
		}
		asset >>= 4
		remain--
	}

	return string(res)
}

// RandLetterNum 生成包含大小写字母和数字的随机字符串。
func RandLetterNum(len int) string {
	res := make([]byte, len)
	asset, remain := rand.Int63(), 10

	for i := 0; i < len; i++ {
		if remain == 0 {
			asset, remain = rand.Int63(), 10
		}
		if bits6 := asset & 0b111111; bits6 < 62 {
			res[i] = numLetters[bits6]
		} else {
			i--
		}
		asset >>= 6
		remain--
	}

	return string(res)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
