package main

import (
	"fmt"

	"github.com/anjanavitthal/go-fibre-crm/database"
	"github.com/anjanavitthal/go-fibre-crm/leads"
	"github.com/gofiber/fiber/v2"
)

func initRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", leads.GetLeads)
	app.Get("/api/v1/leads/:id", leads.GetLeadByID)
	app.Post("/api/v1/leads", leads.NewLead)
	app.Delete("/api/v1/leads/:id", leads.DeleteLead)
}
func main() {

	app := fiber.New()
	database.InitDatabase()
	database.DBConn.AutoMigrate(&leads.Lead{})
	fmt.Print("Database Migrated Successfully")

	initRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()

}
