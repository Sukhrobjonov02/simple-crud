package handlers

import (
	"crud/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	//"github.com/go-playground/validator/v10"
)

func CreateBook(c echo.Context) error {
	book := CreateBookRequest{}
	err := c.Bind(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	req := fromCreateBookRequest(book)

	newBook, err := repositories.CreateBook(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := toCreateBookResponse(newBook)

	return c.JSON(http.StatusCreated, CreateBookResponse(res))
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	book := UpdateBookRequest{}
	err = c.Bind(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedBook, err := repositories.UpdateBook(
		fromUpdateBookRequest(book),
		idInt,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, toUpdateBookResponse(updatedBook))
}

func GetBook(c echo.Context) error {
	id := c.Param(("id"))
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newBook := GetBookRequest{Id: idInt}

	req := fromGetBookRequest(newBook)

	book, err := repositories.GetBook(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := toGetBookResponse(book)

	return c.JSON(http.StatusOK, res)
}

func GetAllBooks(c echo.Context) error {
	books, err := repositories.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}



	return c.JSON(http.StatusOK, toGetAllBooksResponse(books))
}

func DeleteBook(c echo.Context) error {
	id := c.Param(("id"))
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newBook := DeleteBookRequest{Id: idInt}

	req := fromDeleteBookRequest(newBook)

	book, err := repositories.DeleteBook(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := toDeleteBookResponse(book)

	return c.JSON(http.StatusOK, res)
}
