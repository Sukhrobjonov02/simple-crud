package main

import (
	"crud/cmd/handlers"
	"crud/cmd/storage"

	"github.com/labstack/echo/v4"
	//"crud/cmd/repositories"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	storage.InitDB()

	e.POST("/books", handlers.CreateBook)
	e.PUT("/books/:id", handlers.UpdateBook)
	e.GET("/books/:id", handlers.GetBook)
	e.GET("/books", handlers.GetAllBooks)
	e.DELETE("/books/:id", handlers.DeleteBook)

	e.Logger.Fatal(e.Start(":8080"))
}
