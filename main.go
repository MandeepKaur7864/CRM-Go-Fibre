package main

import (
	"fmt"
	"go_crm/database"
	"go_crm/lead"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/getleads", lead.GetLeads)
	app.Get("/api/v1/getlead/:id", lead.GetLead)
	app.Post("/api/v1/createlead", lead.NewLead)
	app.Delete("/api/v1/deletelead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connection opned to database")

	err = database.DBConn.AutoMigrate(&lead.Lead{})
	if err != nil {
		panic("Database migration failed: " + err.Error())
	}
	fmt.Println("Database Migrated!")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	// Properly close the DB connection
	sqlDB, err := database.DBConn.DB()
	if err != nil {
		panic("Failed to get underlying sql.DB: " + err.Error())
	}
	defer sqlDB.Close()
}
