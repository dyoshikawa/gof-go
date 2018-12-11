package main

type Book struct {
	Name string
}

func NewBook(name string) Book {
	book := new(Book)
	book.Name = name
	return *book
}
