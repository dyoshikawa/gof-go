package main

type Book struct {
	Name string
}

func (book *Book) NewBook(name string) {
	book.Name = name
}