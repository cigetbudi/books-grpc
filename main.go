package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/cigetbudi/books-grpc/book/bookpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// struct server buat implementasi
type server struct {
}

func (*server) GetBook(c context.Context, req *bookpb.GetBookRequest) (*bookpb.GetBookResponse, error) {
	//getting request
	bookReq := req.GetBook()

	//retrieving from request
	book := bookpb.Book{
		Id:     bookReq.Id,
		Title:  bookReq.Title,
		Author: bookReq.Author,
		IsRead: false,
	}

	return &bookpb.GetBookResponse{
		Status: true,
		Data:   &book,
	}, nil
}

func main() {
	// ngelog nampilin error baris berapa
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Service Book dijalankan")

	// creating gRPC server
	lis, err := net.Listen("tcp", "0.0.0.0:7676")
	if err != nil {
		log.Fatalf("Gagal listen ke: %v\n", err)
	}

	s := grpc.NewServer()

	//masukin RegisterBookServiceServer ke s
	bookpb.RegisterBoookServiceServer(s, &server{})

	//reflection diaktifin biar bisa ditest
	reflection.Register(s)

	go func() {
		fmt.Println("Menjalankan Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Menunggu ctrl+c biar mati
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Matiin semua kalo udah ctrl+c
	<-ch
	fmt.Println("Memberhentikan server...")
	s.Stop()
	fmt.Println("Memberhentikan listener...")
	lis.Close()
	fmt.Println("Program berhenti")

}
