package handlers

import (
	"crud/cmd/models"
	"crud/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	book := models.Book{}
	err := c.Bind(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newBook, err := repositories.CreateBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	book := models.Book{}

	err = c.Bind(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	updatedBook, err := repositories.UpdateBook(book, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedBook)
}

func GetBook(c echo.Context) error {
	id := c.Param(("id"))
	book := models.Book{}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Bind(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := repositories.GetBook(book, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func GetBooks(c echo.Context) error {
	res, err := repositories.GetBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteBook(c echo.Context) error {
	id := c.Param(("id"))
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := repositories.DeleteBook(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
