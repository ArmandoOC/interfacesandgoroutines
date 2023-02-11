package data

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{ID: 1, Title: "El Oso", Finished: false},
	{ID: 2, Title: "El Perro", Finished: false},
	{ID: 3, Title: "El Gato", Finished: false},
	{ID: 4, Title: "El Pez", Finished: false},
	{ID: 5, Title: "El Rinoceronte", Finished: false},
	{ID: 6, Title: "La yegua", Finished: false},
	{ID: 7, Title: "La hiena", Finished: false},
	{ID: 8, Title: "El tigre", Finished: false},
	{ID: 9, Title: "La hormiga", Finished: false},
	{ID: 10, Title: "El lagarto", Finished: false},
}

func findBook(id int, m *sync.Mutex) (int, *Book) {
	index := -1
	var book *Book

	m.Lock()
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.Unlock()

	return index, book
}

func FinishBook(id int, m *sync.Mutex) {
	i, book := findBook(id, m)
	if i < 0 {
		return
	}

	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()

	fmt.Printf("Finished book: %s", book.Title)
}
