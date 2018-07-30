package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err,ok := r.(error);ok{
			// 接受错误类型
			fmt.Println("Error occuresd:", err)
		}else{
			panic(r)
		}
	}()
	//panic(errors.New("this is an error"))
	panic(123)
}

func main() {
	tryRecover()
}
