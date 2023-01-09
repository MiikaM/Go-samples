package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/miikam/basic-crm/database"
	"github.com/miikam/basic-crm/lead"
)

func main() {
	app := fiber.New()
	initDb()
	setRoutes(app)
	app.Listen(os.Getenv("PORT"))
	defer database.DbConn.Close()
}

func setRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLeadById)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDb() {
	var err error
	database.DbConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Connection opened to database")
	database.DbConn.AutoMigrate(&lead.Lead{})
}
