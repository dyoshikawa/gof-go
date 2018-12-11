package main

type BookShelf struct {
	books []Book
	last  int
}

func NewBookShelf() *BookShelf {
	return new(BookShelf)
}

func (bs *BookShelf) GetBookAt(index int) Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book Book) {
	bs.books = append(bs.books, book)
	bs.last++
}

func (bs *BookShelf) GetLength() int {
	return bs.last
}

func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}
