package main

import (
	"fmt"
	"time"
)


func worker(c chan int){
	for n:= range c {
		fmt.Println(n)
	}
}

func bufferedChannel(){
	c := make(chan int, 3)
	go worker(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	close(c)
	time.Sleep(time.Microsecond)
}

func main() {
	bufferedChannel()
}
