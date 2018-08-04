package main

func main() {
	var c1, c2 chan int

	select {
	case n := <- c1:
		println(n)
	case n := <- c2:
		println(n)
	default:
		println("defalut")
	}
	
}
