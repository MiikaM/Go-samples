package routes

import (
	"github.com/gorilla/mux"
	"github.com/miikam/bookstore-management-api/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
