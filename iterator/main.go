package main

import "fmt"

func main() {
	bs := NewBookShelf()
	bs.AppendBook(NewBook("Book"))
	bs.AppendBook(NewBook("Magazine"))
	it := bs.Iterator()

	for it.HasNext() {
		book := it.Next().(Book)
		fmt.Println(book.Name)
	}
}
