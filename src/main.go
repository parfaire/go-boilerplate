package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// global vars
var username string = "root"
var password string = ""
var address string = "127.0.0.1"
var port string = "3306"
var dbName string = "heroes"
var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, address, port)

func handleRequests() {
	runtime.GOMAXPROCS(12) // 12 threads/childs max
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/heroes/:hero_id", func(c *fiber.Ctx) error {
		hero_id, _ := strconv.Atoi(c.Params("hero_id"))
		hero := select_hero_by_id(hero_id)
		return c.JSON(hero)
	})
	app.Listen(":8080")
}

func main() {
	if !fiber.IsChild() {
		create_db_and_tables(dsn, dbName)
		establish_gorm_connection(dsn, dbName)
		create_heroes()
	} else {
		establish_gorm_connection(dsn, dbName)
	}
	handleRequests()
}
