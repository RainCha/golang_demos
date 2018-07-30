package main

import (
	"fmt"
	"time"
)

func run(task_id, sleeptime int, ch chan string) {

	fmt.Println("start task %d %d", task_id, sleeptime)
	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
	return
}

/**
  利用 time.After 做 goruntine 超时控制
 */
func timeOutRun(task_id, sleeptime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go run(task_id, sleeptime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		ch <- fmt.Sprintf("task id %d , timeout", task_id)
	}
}


func startOutRule(){
	input := []int{3, 2, 1}
	ch := make(chan string)
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		go run(i, sleeptime, ch)
	}

	for range input {
		fmt.Println(<-ch)
	}

	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}

func startInRule(){
	input := []int{3,2,1}
	chs := make([]chan string, len(input))

	startTime := time.Now()

	fmt.Println("Multirun start")

	for i, sleeptime := range input{
		chs[i] = make(chan string)
		go run(i, sleeptime, chs[i])
	}

	for _, ch:=range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}




func startTimeoutRun(){
	input := []int{3, 2, 1}
	// 超时配置
	timeout := 2
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string)
		go timeOutRun(i, sleeptime, timeout, chs[i])
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}

//func main() {
//	// 并发执行， 无序返回
//	startOutRule()
//
//	// 并发执行，有序返回
//	fmt.Println("***************")
//	startInRule()
//
//	// 并发执行，有序返回，超时控制
//	fmt.Println("***************")
//	startTimeoutRun()
//
//	//
//
//
//
//}