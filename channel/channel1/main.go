package main

import (
	"fmt"
	"time"
)


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

func worker(c chan int){
	for n:= range c {
		fmt.Printf("Worker 0 received %d\n", n)
	}
}


func workerId(id int, c chan int){
	for n:= range c {
		fmt.Printf("WorkerId %d received %d\n", id, n)
	}
}


func createWorker(id int) chan<- int{
	c := make(chan int)
	go workerId(id, c)
	return c
}

func channelDemo()  {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}




func main() {
	//bufferedChannel()
	channelDemo()
}
