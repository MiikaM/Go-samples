package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/miikam/bookstore-management-api/pkg/models"
	"github.com/miikam/bookstore-management-api/pkg/utils"
)

var NewBook models.Book

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	resDb, _ := json.Marshal(newBooks)

	res.Header().Set("Content-Type", "application/pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resDb)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	ID := utils.ParseId(req)

	bookDetails, _ := models.GetBookById(ID)
	resDb, _ := json.Marshal(bookDetails)

	res.Header().Set("Content-Type", "application/pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resDb)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)
	book := CreateBook.CreateBook()
	resDb, _ := json.Marshal(book)

	res.Header().Set("Content-Type", "application/pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resDb)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	ID := utils.ParseId(req)

	removedBook := models.DeleteBook(ID)
	resDb, _ := json.Marshal(removedBook)

	res.Header().Set("Content-Type", "application/pkglication/json")
	res.WriteHeader(http.StatusAccepted)
	res.Write(resDb)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(req, updateBook)

	ID := utils.ParseId(req)

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	resDb, _ := json.Marshal(bookDetails)

	res.Header().Set("Content-Type", "application/pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resDb)

}
