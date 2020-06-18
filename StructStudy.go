package main

import "fmt"

type Book struct {
	titile  string
	author  string
	subject string
	book_id int
}

func main() {
	goBook := Book{"Go_语言学习体会", "yxl", "Go_语言", 5201314}
	fmt.Println(goBook)

	goBook2 := Book{titile: "Go_语言学习体会", author: "yxl", book_id: 5201314}
	fmt.Println(goBook2)
	goBook2.subject = "新的语言"
	fmt.Println(goBook2)

	// 使用值传递，在方法内修改结构体变量值得在方法外面是不生效
	detailInfo(goBook2)
	fmt.Println(goBook2)
	// 使用指针，则是引用传递，在方法内修改结构体变量值可以生效
	updateBook(&goBook2)
	fmt.Println(goBook2)

}


func detailInfo(book Book){
	book.book_id = 1314520
	fmt.Printf("book titile : %s \n",book.titile)
	fmt.Printf("book author : %s\n",book.author)
	fmt.Printf("book subject : %s\n",book.subject)
	fmt.Printf("book id : %d\n",book.book_id)
}

func updateBook(pbook *Book){
	pbook.book_id = 1314520
	fmt.Printf("book titile : %s \n",pbook.titile)
	fmt.Printf("book author : %s\n",pbook.author)
	fmt.Printf("book subject : %s\n",pbook.subject)
	fmt.Printf("book id : %d\n",pbook.book_id)
}