package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/miikam/basic-crm/database"
)

type Lead struct {
	gorm.Model

	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DbConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLeadById(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DbConn

	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DbConn
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DbConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(404).Send("No lead found with ID")
		return
	}

	db.Delete(&lead)
	c.Send("Lead deleted successfully")
}
