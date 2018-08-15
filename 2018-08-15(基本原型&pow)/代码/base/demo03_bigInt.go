package main

import (
	"fmt"
	"math/big"
)

func main() {
	/*
	fibonacci数列


	big包：中可以操作取值范围较大的数据：Int，Float。。。
		big.Int --struct：类

		step1：创建big.Int对象：指针对象
			内置函数new()
			new(big.Int)-->bi1

			big包下提供的：
			big.NewInt(int64)-->bi2


		step2：设置数值
			setXXX()
				setString()
				setInt()
				setBytes()...

		step3：运算操作：
			调用对应的方法：Add(),Sub(),Div(),Mul(),Lsh(),Rsh().....

	 */
	//fmt.Println(fibo(100))
	for i := 1; i <= 1000; i++ {
		//fmt.Println(i, fibo(i))
		fmt.Println(i, fibo2(i))
	}

	//创建Int类的对象，用于存储大证书
	var bi1 big.Int
	fmt.Println(bi1)
	bi1.SetInt64(100)
	fmt.Println(bi1)
	fmt.Printf("%T\n", bi1) //big.Int

	bi2 := big.NewInt(2) //pointer
	fmt.Println(bi2)
	fmt.Printf("%T\n", bi2)

	bi3 := new(big.Int)
	fmt.Printf("%T\n", bi3)
	bi3.SetString("5", 10)
	fmt.Println(bi3)

	//bi3.SetBytes([]byte{1, 2, 3}) // [0,0,0,0, 0,1,2,3] // 255
	//fmt.Println(bi3)

	fmt.Println(1*255*255 + 2*255 + 3*1) //?

	//操作运算：方法
	fmt.Println(bi2, bi3)

	//func (z *Int) Add(x, y *Int) *Int
	bi4 := new(big.Int)
	//bi4 = bi4.Add(bi2,bi3)
	bi4.Add(bi2, bi3)
	fmt.Println(bi2, bi3, bi4)

	bi5 := new(big.Int)
	bi5.Div(bi3, bi2)
	fmt.Println(bi2, bi3, bi5)

	//左移：10 << 2 ---> 40
	fmt.Println(10 << 2)

	bi6 := big.NewInt(10)

	bi6.Lsh(bi6, 2)
	fmt.Println(bi6)

}

var arr [1001]int

func fibo(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = arr[n-1] + arr[n-2]
	}
	arr[n] = res
	return
}

var arr2 [1001]*big.Int

func fibo2(n int) (res *big.Int) {
	if n <= 2 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(arr2[n-1], arr2[n-2])
	}
	arr2[n] = res
	return
}
