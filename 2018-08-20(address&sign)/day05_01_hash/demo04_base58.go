package main

import (
	"math/big"
	"bytes"
	"fmt"
)

func main() {
	/*
	测试Base58
	 */

	//1.原始数据
	data1 := []byte{10, 20, 30, 40, 50, 60}
	data1 = append([]byte{0},data1...) //前缀+数据
	fmt.Println("原始数据：",data1)
	//Base58编码
	encode:=Base58Encode(data1)
	fmt.Println("Base58编码后：",encode) //[]byte
	fmt.Println(string(encode))
	//Base58解码
	decode:=Base58Decode(encode)
	fmt.Println("Base58解码后：",decode)


	//测试数据
	data2:="wangergou"
	encode2:=Base58Encode([]byte(data2))
	fmt.Println(encode2)
	fmt.Println(string(encode2))

	decode2:=Base58Decode(encode2)
	fmt.Println(decode2)
	fmt.Println(string(decode2))
	fmt.Println(string(decode2[1:])) //wangergou
}

//base64

/*
ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/
0(零)，O(大写的o)，I(大写的i)，l(小写的L),+,/
 */

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

//字节数组转Base58，加密
func Base58Encode(input []byte) []byte {
	var result [] byte
	x := big.NewInt(0).SetBytes(input)
	//fmt.Println("x:",x)
	base := big.NewInt(int64(len(b58Alphabet)))//58
	//fmt.Println("base:",base)
	zero := big.NewInt(0)
	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		//fmt.Println("x:",x,",mod:",mod) //0-57
		result = append(result, b58Alphabet[mod.Int64()])
		//fmt.Println("result:",string(result))
	}
	//将得到result中的数据进行反转
	ReverseBytes(result)
	for b := range input { //遍历input数组：index,value
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}
	//1 result

	return result

}

//Base58转字节数组，解密
func Base58Decode(input [] byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}
	fmt.Println("zeroBytes:",zeroBytes)
	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded
}

//字节数组反转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
