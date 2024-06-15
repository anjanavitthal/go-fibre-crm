package leads

import (
	"github.com/anjanavitthal/go-fibre-crm/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Company string `json:"company"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn

	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLeadByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	var lead = new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).SendString("err")
		return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.Params("id")
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found")

	}

	db.Delete(&lead)
	return c.SendString(" lead successfuly seleted")
}
