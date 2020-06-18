package main

import "fmt"

func main() {
	//var a string = "hello,jack!"
	//fmt.Println(a)
	//var i, j int = 1, 2
	//fmt.Println(i + j)
	//var b bool = true
	//fmt.Println(b)
	//c := false;
	//fmt.Println(c)
	//data := 520
	//fmt.Println(data)
	//b,intValue ,boolValue, strValue := false,234,false,"jack,hello"
	//
	//if(!boolValue){
	//	fmt.Println(strValue,intValue,b)
	//}

	a, b := 1314, 520
	//maxValue := max(a, b)
	//fmt.Println(maxValue)
	//oldAValue, oldBValue := swap(a, b)
	//fmt.Print(oldAValue, oldBValue)

	fmt.Printf("before swap a = %d b = %d \n", a, b)
	swap2(&a,&b);
	fmt.Printf("after swap a = %d b = %d \n", a, b)

	// 定义并初始化数组
	var balance = [...]float32{10000,2,3,4,7,50}
	balance[4] = 50.0

}

func max(a int, b int) int {
	if (a > b) {
		return a
	} else {
		return b
	}
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swap2(a *int, b *int) {
	temp := *a;
	*a = *b;
	*b = temp;
}
