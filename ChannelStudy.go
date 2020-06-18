package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{0, 2, 4, 6, 8}

	c := make(chan int)
	c2 := make(chan int)
	go sum(arr1, c)
	go sum(arr2, c2)
	x, y := <-c, <-c2 // 从通道 c 中接收

	fmt.Printf("%d + %d = %d\n",x, y, x+y)
}
