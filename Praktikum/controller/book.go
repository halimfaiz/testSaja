package controller

import (
	"Praktikum/model"
	"Praktikum/repository/database"
	"Praktikum/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController interface {
	GetBooksController(c echo.Context) error
	GetBookController(c echo.Context) error
	CreateBookController(c echo.Context) error
	DeleteBookController(c echo.Context) error
	UpdateBookController(c echo.Context) error
}

type bookController struct {
	bookUsecase    usecase.BookUsecase
	bookRepository database.BookRepository
}

func NewBookController(
	bookUsecase usecase.BookUsecase,
	bookRepository database.BookRepository,
) *bookController {
	return &bookController{
		bookUsecase,
		bookRepository,
	}
}

// get all Books
func (b *bookController) GetBooksController(c echo.Context) error {
	books, e := b.bookRepository.GetBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func (b *bookController) GetBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := b.bookRepository.GetBook(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages":         "Error Get Book",
			"ErrorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    book,
	})
}

// create new book
func (b *bookController) CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := b.bookUsecase.CreateBook(&book); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "error create  book",
			"book":    err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func (b *bookController) DeleteBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := b.bookUsecase.DeleteBook(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error Delete book",
			"errorDescription": err,
			"errorMessage":     "buku tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Delete book",
	})
}

// update book by id
func (b *bookController) UpdateBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book := model.Book{}
	c.Bind(&book)
	book.ID = uint(id)

	if err := b.bookUsecase.UpdateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error update book",
			"errorDescription": err,
			"errorMessage":     "buku tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}
