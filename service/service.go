package service

import (
	"github.com/cigetbudi/books-grpc/model"
	"github.com/cigetbudi/books-grpc/repository"
)

func AddBook(bookData model.Book) model.Book {
	return repository.AddBook(bookData)
}

func GetBook(bookId string) (int, model.Book) {
	return repository.GetBook(bookId)
}

func UpdateBook(bookData model.Book, id string) model.Book {
	return repository.UpdateBook(bookData, id)
}

func DeleteBook(id string) bool {
	return repository.DeleteBook(id)
}
