package utils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

// Random
// @description: 随机数工具接口
type Random interface {
	GetRandomInt(len int) string         // 获取指定长度的随机数字字符串
	GetRandomString(len int) string      // 生成随机字符串
	GetRandomStringLower(len int) string // 生成小写随机字符串
	GetRandomStringMini(len int) string  // 获取去掉了iI0O1的随机字符串
	GetRandomNumber(min, max int) int    // 获取指定范围内的一个随机数
}

type random struct{}

// RandomUtils
// @description: 随机数工具
// @return Random
func RandomUtils() Random {
	return &random{}
}

// GetRandomInt
// @description: 获取指定长度的随机数字字符串
// @receiver random
// @param len
// @return string
func (random) GetRandomInt(len int) string {
	var numbers = []byte{0, 1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= len; i++ {
		rd, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		container += fmt.Sprintf("%d", numbers[rd.Int64()])
	}
	return container
}

// GetRandomString
// @description: 生成随机字符串
// @receiver random
// @param len
// @return string
func (random) GetRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// GetRandomStringLower
// @description: 生成小写的随机字符串
// @receiver r
// @param len
// @return string
func (r random) GetRandomStringLower(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyz1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// GetRandomStringMini
// @description: 获取去掉了iI0O1的随机字符串
// @receiver random
// @param len
// @return string
func (random) GetRandomStringMini(len int) string {
	var container string
	var str = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// GetRandomNumber
// @description: 获取指定范围内的一个随机数
// @receiver random
// @param min
// @param max
// @return int
func (random) GetRandomNumber(min, max int) int {
	en, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	rn := int(en.Int64()) + min
	return rn
}
