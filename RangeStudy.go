package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0;
	for index, num := range nums {
		fmt.Printf("nums[%d] = %d \n", index, num)
		sum += num
	}
	fmt.Printf("sum is %d \n", sum)

	maps := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range maps {
		fmt.Printf("%s ---> %s \n", k, v)
	}

}
