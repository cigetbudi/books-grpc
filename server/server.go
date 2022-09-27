package server

import (
	"context"
	"errors"

	"github.com/cigetbudi/books-grpc/book/bookpb"
	"github.com/cigetbudi/books-grpc/model"
	"github.com/cigetbudi/books-grpc/service"
)

// Server GRPC struct
type Server struct {
}

// mapping book model ke bookpb Book
func mapToBookpb(book model.Book) *bookpb.Book {
	return &bookpb.Book{
		Id:     book.Id,
		Author: book.Author,
		Title:  book.Title,
		IsRead: book.IsRead,
	}
}

// SEMUA FUNGSI DI SINI BASEON BookServiceServer interface

func (*Server) GetBook(c context.Context, req *bookpb.GetBookRequest) (*bookpb.GetBookResponse, error) {
	// Ambil id dari Vm req
	bookId := req.GetId()

	// masukin id ke service getbook
	_, res := service.GetBook(bookId)

	// kalo kosong, return error dan ngisi status sama nil ke model reponse
	if res.Id != bookId {
		return &bookpb.GetBookResponse{Status: false, Data: nil}, errors.New("Data tidak dapat ditemukan")
	}

	// kalo ada, hasil getbook dimasukin ke object getbook reponse
	var bookData *bookpb.Book = &bookpb.Book{
		Id:     res.Id,
		Author: res.Author,
		Title:  res.Title,
		IsRead: res.IsRead,
	}

	return &bookpb.GetBookResponse{
		Status: true,
		Data:   bookData,
	}, nil
}

func (*Server) AddBook(c context.Context, req *bookpb.AddBookRequest) (*bookpb.AddBookResponse, error) {

	// menerima request
	var bookReq *bookpb.Book = req.GetBook()

	// masukin bookData ke storage / db sementara array
	var bookData model.Book = model.Book{
		Id:     bookReq.GetId(),
		Title:  bookReq.GetTitle(),
		Author: bookReq.GetTitle(),
		IsRead: bookReq.GetIsRead(),
	}

	//panggil service addbook buat nambah data buku
	var newBook model.Book = service.AddBook(bookData)

	//nampilin yang baru aja ditambahin
	return &bookpb.AddBookResponse{
		Status: true,
		Data:   mapToBookpb(newBook),
	}, nil
}

func (*Server) UpdateBook(c context.Context, req *bookpb.UpdateBookRequest) (*bookpb.UpdateBookResponse, error) {

	// menerima request semacam c.bind(vm)
	var bookReq *bookpb.Book = req.GetBook()

	// Variabel bookData untuk mengubah buku
	var bookData model.Book = model.Book{
		Id:     bookReq.GetId(),
		Title:  bookReq.GetTitle(),
		Author: bookReq.GetAuthor(),
		IsRead: bookReq.GetIsRead(),
	}

	// panggil service UpdateBook, masukin nilai baru ke nilai lama  berdasarkan IDNya berapa
	var updatedBook model.Book = service.UpdateBook(bookData, bookData.Id)

	// resonsenya nilai yang udah dirubah
	return &bookpb.UpdateBookResponse{
		Status: true,
		Data:   mapToBookpb(updatedBook),
	}, nil
}

func (*Server) DeleteBook(c context.Context, req *bookpb.DeleteBookRequest) (*bookpb.DeleteBookResponse, error) {
	// get id Dari request
	bookId := req.GetBookId()

	//panggil service deletebook by ID
	res := service.DeleteBook(bookId)

	return &bookpb.DeleteBookResponse{
		Status: res,
	}, nil
}
