package main

import (
	"os"
	"bufio"
	"golang_demos/defer/fib"
	"fmt"
)

func tryDefer() {
	//defer println(1)
	//defer println(2)
	//println(3)

	// defer 参数在defer 语句时计算
	for i := 0; i <= 30; i++ {
		defer println(i)
		if i == 30 {
			panic(error.Error)
		}
	}
}

func writeFile(filename string)  {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// 使用 buffer 写入文件更快？？
	writer := bufio.NewWriter(file)

	// 不写 flush 没有内容？？
	defer writer.Flush()

	f := fib.Fib()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	tryDefer()

	//writeFile("tx.txt")
}

