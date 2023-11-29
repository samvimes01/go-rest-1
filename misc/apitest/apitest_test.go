package apitest

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

// create a test case for getting all books data
func TestGetBooks_Success(t *testing.T) {
	// create a test
	apitest.New().
		// add a handler that will be tested
		HandlerFunc(getAllBooks).
		// send a GET request to get all books data
		Get("/books").
		// the expected response status code is 200
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreateBook_Success(t *testing.T) {
	// create a request body to create a new book
	var newBook Book = Book{
		Id:    1,
		Title: "test book",
	}

	// create a test
	apitest.New().
		// add a handler that will be tested
		HandlerFunc(createBook).
		// send a POST request to create a new book
		Post("/books/create").
		// set the request body in JSON format
		JSON(newBook).
		// the expected response status code is 201
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateBook_Failed(t *testing.T) {
	// create a request body to create a new book
	var newBook Book = Book{
		Id:    1,
		Title: "test book",
	}

	// create a test
	apitest.New().
		// add a handler that will be tested
		HandlerFunc(createBook).
		// send a GET request to create a new book
		Get("/books/create").
		// set the request body in JSON format
		JSON(newBook).
		// the expected response status code is 405
		Expect(t).
		Status(http.StatusMethodNotAllowed).
		End()
}