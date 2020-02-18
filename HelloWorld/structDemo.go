package main

import "fmt"

func main()  {

	//创建一个新的结构体
	fmt.Println(Books{"Go 语言", "WD", "Go 语言教程", 6495407})

	var book1 Books
	var book2 Books

	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	book1.book_id = 6495407

	book2.title = "Python 教程"
	book2.author = "www.runoob.com"
	book2.subject = "Python 语言教程"
	book2.book_id = 6495700

	fmt.Printf("book 1 \n title: %s\n author: %s\n subject: %s\n book_id: %d\n", book1.title, book1.author,
		book1.subject, book1.book_id)
	fmt.Printf("book 2 \n title: %s\n author: %s\n subject: %s\n book_id: %d\n", book2.title, book2.author,
		book2.subject, book2.book_id)

	//结构体作为函数参数
	//printBook(book1)
	//printBook(book2)

	//结构体指针
	printBook(&book1)
}

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func printBook( book *Books)  {
	fmt.Printf("book \n title: %s\n author: %s\n subject: %s\n book_id: %d\n", book.title, book.author,
		book.subject, book.book_id)
}