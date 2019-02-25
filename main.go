package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang pointers",
		Author: "Mr.Goland", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr.Goroutine", Year: "2012"},
		Book{ID: 4, Title: "Golang concurrency", Author: "Mr.Goroutine", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr.Goroutine", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(reflect.TypeOf(params["id"]))
	i, _ := strconv.Atoi(params["id"])
	log.Println(reflect.TypeOf(i))

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}
func addBooks(w http.ResponseWriter, r *http.Request) {
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("update a book is called")
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a book is called")
}
