package main

type BookShelf struct {
	books []Book
	last int
}

func NewBookShelf(maxSize int) *BookShelf {
	bs := new(BookShelf)
	return bs
}

func (bs *BookShelf) GetBookAt(index int) Book {
	return bs.books[index]
}

func (bs *BookShelf) appendBook(book Book) {
	bs.books = append(bs.books, book)
	bs.last++
}

func (bs *BookShelf) GetLength() int {
	return bs.last
}

func (bs *BookShelf) Iterator() *BookShelfIterator {
	return NewBookShelfIterator(*bs)
}