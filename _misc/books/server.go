package books

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// create a book type
type Book struct {
	Id    int
	Title string
}

// create a books data
var books []Book = []Book{
	{
		Id:    1,
		Title: "Programming with Go",
	},
	{
		Id:    2,
		Title: "Introduction to MySQL",
	},
	{
		Id:    3,
		Title: "Microservice Architecture",
	},
}

func SrvBooks() {
	// create a mux to handle the requests
	var mux *http.ServeMux = http.NewServeMux()

	// register handlers to mux
	mux.HandleFunc("/books", getAllBooks)
	mux.HandleFunc("/books/create", createBook)

	// create the server with the mux as a handler
	var server *http.Server = &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	fmt.Println("Starting server on port 3000")

	// start the server
	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

// getAllBooks returns all books data
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	// convert "books" into JSON format
	output, err := json.Marshal(books)
	if err != nil {
		log.Panic(err)
	}

	// set the response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// set the response body
	_, err = w.Write(output)
	if err != nil {
		log.Panic(err)
	}
}

// createBook returns the recently created book
func createBook(w http.ResponseWriter, r *http.Request) {
	// check the request method
	if r.Method != http.MethodPost {
		// if the request method is not "POST"
		// this response is returned
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("method not allowed"))
		if err != nil {
			log.Panic(err)
		}
		return
	}

	// create a variable to store the request body for book data
	var newBook Book

	// decode the request body
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		log.Panic(err)
	}

	// add a new book
	books = append(books, newBook)

	// convert the recently added book data into the JSON format
	output, err := json.Marshal(newBook)
	if err != nil {
		log.Panic(err)
	}

	// set the response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// set the response body
	_, err = w.Write(output)
	if err != nil {
		log.Panic(err)
	}
}