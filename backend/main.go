package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// Create product struct
	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}

	// Connect to db
	host := os.Getenv("DB_SERVER")
	DBDriver := fmt.Sprintf("host=%s port=5432 user=root dbname=base password=root", host)
	log.Println("start ---")
	log.Println(DBDriver)
	log.Println("--- end")
	// db, err := gorm.Open("postgres", DBDriver)
	// if err != nil {
	// 	panic("failed to connect db")
	// 	log.Fatalln("failed to connect db - log")
	// }
	// log.Println("success db connection")
	// defer db.Close()

	// Create new products
	// After db connection is created.
	// db.CreateTable(&Product{})

	// Also some useful functions
	// db.HasTable(&Product{})          // =>;; true
	// db.DropTableIfExists(&Product{}) //Drops the table if already exists

	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	// db.First(&product, 1) // find product with id 1

	fmt.Println(product)

	// Run backend app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, from fiber")
	})

	app.Listen(3000)
}
