package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

var books = []Book{
	{ID: 1, Title: "Go Programming", Author: "John Doe"},
	{ID: 2, Title: "Kubernetes", Author: "Kelsey Hightower"},
	{ID: 3, Title: "Docker", Author: "Nigel Poulton"},
}

func main() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookHandler)

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
		return
	}

	if r.Method == http.MethodPost {
		var book Book

		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		books = append(books, book)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]

	if r.Method == http.MethodGet {
		for _, book := range books {
			if fmt.Sprint(book.ID) == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(book)
				return
			}
		}

		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodPut {
		var updatedBook Book

		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		for i, book := range books {
			if fmt.Sprint(book.ID) == id {
				updatedBook.ID = book.ID
				books[i] = updatedBook

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(updatedBook)
				return
			}
		}

		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodDelete {
		for i, book := range books {
			if fmt.Sprint(book.ID) == id {
				books = append(books[:i], books[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		http.NotFound(w, r)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
