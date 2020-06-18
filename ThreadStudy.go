package main

import (
	"fmt"
	"time"
)

func say(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

/**
 * 可通过go 关键字来开启goroutine 并发，goroutine 是轻量级线程，goroutine 的调度室由Golang运行时进行管理的
 */
func main() {
	go say("jack")
	say("yxl")

}
