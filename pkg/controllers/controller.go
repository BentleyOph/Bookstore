package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/BentleyOph/bookstore/pkg/models"
	"github.com/BentleyOph/bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBooksById retrieves book details by ID.
// It takes a http.ResponseWriter and a http.Request as parameters.
// It parses the book ID from the request URL and calls the GetBooksById function from the models package to fetch the book details.
// The retrieved book details are then marshaled into JSON format and written to the response writer.
// If there is an error while parsing the book ID, an error message is printed to the console.
// The response writer is set with the appropriate content type and status code before writing the response.

func GetBooksById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBooksById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatebook = &models.Book{}
	utils.ParseBody(r, updatebook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal("Error while parsing")
	}
	bookdetails, db := models.GetBooksById(ID)
	if updatebook.Name != "" {
		bookdetails.Name = updatebook.Name
	}
	if updatebook.Author != "" {
		bookdetails.Author = updatebook.Author
	}
	if updatebook.Publication != "" {
		bookdetails.Publication = updatebook.Publication
	}
	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set(("Content-Type"), "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
