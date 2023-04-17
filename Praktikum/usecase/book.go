package usecase

import (
	"Praktikum/model"
	"Praktikum/repository/database"
	"errors"
	"fmt"
)

type BookUsecase interface {
	CreateBook(book *model.Book) error
	GetBook(id uint) (book model.Book, err error)
	GetListBooks() (books []model.Book, err error)
	UpdateBook(book *model.Book) (err error)
	DeleteBook(id uint) (err error)
}

type bookUsecase struct {
	bookRepository database.BookRepository
}

func NewBookUsecase(bookRepo database.BookRepository) *bookUsecase {
	return &bookUsecase{bookRepository: bookRepo}
}

func (b *bookUsecase) CreateBook(book *model.Book) error {

	if book.Judul == "" {
		return errors.New("Judul Buku tidak boleh kosong")
	}
	if book.Penulis == "" {
		return errors.New("Penulis Buku tidak boleh kosong")
	}

	err := b.bookRepository.CreateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (b *bookUsecase) GetBook(id uint) (book model.Book, err error) {
	book, err = b.bookRepository.GetBook(id)
	if err != nil {
		fmt.Println("GetBook :Error Getting book from repository")
		return
	}
	return
}

func (b *bookUsecase) GetListBooks() (books []model.Book, err error) {
	books, err = b.bookRepository.GetBooks()
	if err != nil {
		fmt.Println("GetListBooks : Error Getting book from repository")
		return
	}
	return
}

func (b *bookUsecase) UpdateBook(book *model.Book) (err error) {
	err = b.bookRepository.UpdateBook(book)
	if err != nil {
		fmt.Println("UpdateBook : Error updating Book, err: ", err)
		return
	}
	return
}

func (b *bookUsecase) DeleteBook(id uint) (err error) {
	book := model.Book{}
	book.ID = id
	err = b.bookRepository.DeleteBook(&book)
	if err != nil {
		fmt.Println("DeleteBook : Error deleting Book, err: ", err)
		return
	}
	return
}
