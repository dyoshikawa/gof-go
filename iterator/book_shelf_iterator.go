package main

type BookShelfIterator struct {
	BookShelf BookShelf
	index int
}

func NewBookShelfIterator(bs BookShelf) *BookShelfIterator {
	bsi := new(BookShelfIterator)
	bsi.BookShelf = bs
	bsi.index = 0
	return bsi
}

func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.BookShelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (bsi *BookShelfIterator) Next() Book {
	book := bsi.BookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}