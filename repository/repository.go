package repository

import "github.com/cigetbudi/books-grpc/model"

// Database lokal seb[mentara pake slice
var storage []model.Book = []model.Book{}

// Tambah Data buku
func AddBook(bookData model.Book) model.Book {
	storage = append(storage, bookData)
	return bookData
}

// Get buku by id
func GetBook(bookId string) (int, model.Book) {
	for index, v := range storage {
		if v.Id == bookId {
			return index, v
		}
	}
	return 0, model.Book{}
}

// Edit data buku
func UpdateBook(bookVM model.Book, id string) model.Book {
	index, book := GetBook(id)

	book.Title = bookVM.Title
	book.Author = bookVM.Author
	book.IsRead = bookVM.IsRead
	storage[index] = book
	return book
}

// Menghapus data buku by id
func DeleteBook(id string) bool {
	var afterDelete []model.Book = []model.Book{}
	for _, v := range storage {
		if id != v.Id {
			afterDelete = append(afterDelete, v)
		}
	}
	storage = afterDelete
	return true
}
