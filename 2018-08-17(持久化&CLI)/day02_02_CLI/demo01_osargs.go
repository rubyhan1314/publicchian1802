package main

import (
	"os"
	"fmt"
)

func main() {
	/*
	os包：
		控制台的参数：
		os.Args,[]string,一个切片用于存储终端的参数。第一个参数，永远时当前程序的路径。


	 */

	 argsAll:=os.Args
	 //fmt.Printf("%T\n",argsAll)
	 fmt.Println("所有的参数：",argsAll)
	 //获取可用的参数
	 argsUseful:=os.Args[1:]
	 fmt.Println("可用的参数：",argsUseful)
}
/*
go run main.go

go run main.go aaaa bbb  ccc hello

go build -o aa demo01_osargs.go---->aa

./aa aaaa bbbb cccc dddd
 */