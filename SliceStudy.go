package main

import "fmt"

/**
 * 切面相关的知识，可动态变化长度的数组
 */
func main() {

	sliceTest := make([]int, 3)
	sliceTest[0] = 1
	sliceTest[1] = 2
	sliceTest[2] = 3
	fmt.Println(sliceTest)

	ints := append(sliceTest, 4)
	fmt.Println(ints)

	// 查看ints的容量大小
	fmt.Printf("ints cap is : %d\n", cap(ints))

	// 创建切片，并通过copy函数复制
	copyArr := make([]int, len(sliceTest))
	copy(copyArr, sliceTest)

	fmt.Printf("copyArr length is %d\n",len(copyArr))

	for i := 0; i < len(copyArr); i++ {
		fmt.Printf("copyArr[%d] = %d\n",i,copyArr[i]);
	}

	// 截取切片
	fmt.Println("sliceTest[1:3] == ",sliceTest[1:3])

}
