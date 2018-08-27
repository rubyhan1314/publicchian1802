package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {
	/*
	step1：下载otto，虚拟机的代码(go语言写的虚拟机)
		打开终端：输入命令
			go get github.com/robertkrimen/otto

	step2:下载依赖包
		cd $GOPATH/src
		go get gopkg.in/sourcemap.v1


	step3：创建工程
		otto1802
			注意点：
				工程名字：otto1802
				github.com
				gopkg.in，
				并列在src目录下

	 */
	 //测试虚拟机能够运行js
	 http.HandleFunc("/",sayHttp)
	 err:=http.ListenAndServe(":7777",nil)
	 if err != nil{
	 	log.Panic("ListenAndServe:",err)
	 }
}

func sayHttp(w http.ResponseWriter,r *http.Request){
	parseJSFunc(w,r)
}

// 解析 js function
func parseJSFunc(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	//fmt.Printf("%T\n",r)
	//fmt.Printf("%T\n",r.Form)
	//fmt.Printf("%v\n",r.Form)
	for k,v:=range r.Form{
		if k == "jstest"{
			//创建虚拟机
			vm:=otto.New()
			var src = string(v[0])
			//设置web中传递的js代码， 并在虚拟机上运行js
			vm.Run(src)

			//调用js对象
			var obj,_ = vm.Object("object1")
			//设置obj中的值
			obj.Set("age",1000)

			//调用obj中地址
			var name,_=obj.Get("name")
			var age,_=obj.Get("age")
			fmt.Println("name:",name,",age:",age)

			//调用对象的方法
			var meth,_=obj.Call("sayHi",nil,nil)
			fmt.Println(meth)

			var mpar,_=obj.Call("sayHello",11,20)
			fmt.Println(mpar)

		}
	}


	fmt.Fprintf(w,"Hello World!")
}
