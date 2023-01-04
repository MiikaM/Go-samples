package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/miikam/bookstore-management-api/pkg/routes"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Starting movie service at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
