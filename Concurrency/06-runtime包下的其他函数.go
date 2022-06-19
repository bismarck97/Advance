package main

import (
	"fmt"
	"runtime"
)

func main() {
	//GOROOT返回Go的根目录。如果存在GOROOT环境变量，返回该变量的值；否则，返回创建Go时的根目录。
	fmt.Println(runtime.GOROOT())
	//返回Go的版本字符串。它要么是递交的hash和创建时的日期；要么是发行标签如"go1.3"。
	fmt.Println(runtime.Version())
	//NumCPU返回本地机器的逻辑CPU个数。
	fmt.Println(runtime.NumCPU())
}
