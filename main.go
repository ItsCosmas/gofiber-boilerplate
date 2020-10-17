package main

import (
	app "github.com/ItsCosmas/gofiber-boilerplate/api"
	_ "github.com/ItsCosmas/gofiber-boilerplate/api/docs" // Swagger Docs
)

func main() {
	app.Run()
}
