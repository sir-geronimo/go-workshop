package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sir-geronimo/go-workshop/pkg/models"
	"github.com/sir-geronimo/go-workshop/pkg/utils"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)

	b := book.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, _ *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	b := &models.Book{}
	utils.ParseBody(r, b)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book, db := models.GetBookById(ID)
	if b.Name != "" {
		book.Name = b.Name
	}
	if b.Author != "" {
		book.Author = b.Author
	}
	if b.Publication != "" {
		book.Publication = b.Publication
	}
	db.Save(&book)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
