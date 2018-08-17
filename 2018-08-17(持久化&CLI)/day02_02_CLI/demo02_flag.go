package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main() {
	/*
	flag包
		是go中用于解析命令行参数的包。
			程序中可以根据参数的不同，执行不同的逻辑部分。


	第一个部分：解析命令
		step1：定义FlagSet命令对象
		flagCmd:=flag.NewFlagSet("命令名",错误处理)

		step2：解析命令
		flagCmd.Parse(os.Args[2:])

		step3：判断命令是否被解析
			if flagCmd .Parsed(){
				执行该命令对应的操作
			}


	第二个部分：命令后的参数
		获取命令后的参数：

		func String(name string, value string, usage string) *string
		func Int()
		func Bool()
		第一个参数：对应的标签的名
		第二个参数：默认值
		第三个参数：说明


		stringData:=flagCmd.String("flagname","","")



	 */
	 //1.定义接受命令行参数的指针
	 stringData:=flag.String("username","admin","用户名")
	 intData:=flag.Int("age",6,"年龄")
	 //特殊注意点：终端输入-open，就表示true
	 boolData:=flag.Bool("open",false,"是否打开")
	 //fmt.Printf("%T\n",stringData) //*string

	 //2.解析,之后就可以获取数据
	 flag.Parse()

	 //3.获取标签后的数据
	 fmt.Printf("%T,字符串的数据：%s\n",stringData,*stringData)
	 fmt.Printf("%T,整数的数据：%d\n",intData,*intData)
	 fmt.Printf("%T,bool的数据：%t\n",boolData,*boolData)



	 fmt.Println("-------------------------------------------")
	 //1.创建一个命令的标签
	 flagUsernameCmd:=flag.NewFlagSet("input",flag.ExitOnError)
	 //fmt.Printf("%T\n",flagUsernameCmd) //*flag.FlagSet
	 nameData:=flagUsernameCmd.String("name","","姓名")


	 //2.解析该命令
	 if os.Args[1] == "input"{
	 	//fmt.Println("input........")
	 	//解析name后的数据
	 	err:=flagUsernameCmd.Parse(os.Args[2:])
	 	if err != nil{
	 		log.Panic(err)
		}
	 }

	 if flagUsernameCmd.Parsed(){
	 	fmt.Println("name:",*nameData)
	 }




}
/*
右键执行：go run main.go

终端中：
go run demo02.go -username wangergou
go run demo02.go -age 30

或者编译
go build -o bb demo02.go
./bb -age 30



./bb input

	os.Args[1]


./bb input -name wangergou

	input命令：wangergou

	os.Args[0],程序的路径
	os.Args[1],命令：input
	os.Args[2:],-name wangergou


 */