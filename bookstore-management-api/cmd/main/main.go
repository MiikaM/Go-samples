package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/miikam/bookstore-management-api/pkg/routes"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Starting movie service at port <>")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
