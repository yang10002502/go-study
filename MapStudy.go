package main

import "fmt"

func main() {
	var maps map[string]string
	maps = make(map[string]string)
	maps["France"] = "巴黎"
	maps["Italy"] = "罗马"
	maps["Japan"] = "东京"
	maps["India"] = "新德里"

	for county := range maps {
		fmt.Println(county, "首都是 ", maps[county])
	}

	countyName := "Japan"
	// 删除map的元素
	delete(maps, countyName)

	fmt.Println("-----------------------------------------")
	fmt.Printf("删除 %s 国家之后\n",countyName)

	for county := range maps {
		fmt.Println(county, "首都是 ", maps[county])
	}

	fmt.Println("-----------------------------------------")
	countyName = "American"
	// 查看map中是否包含该元素
	capital, ok := maps[countyName]
	if (ok) {
		fmt.Println(countyName, "首都是 ", capital)
	} else {
		fmt.Println("集合中不包含", countyName)
	}

	fmt.Println("-----------------------------------------")
	countyName = "Italy"
	capital, ok = maps[countyName]
	if (ok) {
		fmt.Println(countyName, "首都是 ", capital)
	} else {
		fmt.Println("集合中不包含", countyName)
	}

}
