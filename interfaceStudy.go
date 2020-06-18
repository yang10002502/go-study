package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I'm Nokia Phone ,I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I'm iPhone ,I can call you!")
}

func main() {
	// 声明并初始化
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	// 重新赋值
	phone = new(IPhone)
	phone.call()

}
