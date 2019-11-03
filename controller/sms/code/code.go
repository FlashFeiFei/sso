package code

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

//获取存数字的验证码
func GetRandNumberCode(code_len int) (string) {
	if code_len <= 0 {
		code_len = 4
	}

	code := make([]string, code_len)
	var rand_number *big.Int
	rand_number = big.NewInt(10)

	for i := range code {
		my_rand, _ := rand.Int(rand.Reader, rand_number)
		code[i] = strconv.FormatInt(my_rand.Int64(), 10)
	}

	return strings.Join(code, "")
}

//随机获取字符串加数字的数组随机数
func GetRandCodeByStrAndNumber(code_len int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	code := make([]rune, code_len)
	var rand_number *big.Int
	rand_number = big.NewInt(int64(len(letterRunes)))
	for i := range code {
		my_rand, _ := rand.Int(rand.Reader, rand_number)
		code[i] = letterRunes[my_rand.Int64()]
	}
	return string(code)
}
