package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/cigetbudi/books-grpc/book/bookpb"
	"github.com/cigetbudi/books-grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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
	bookpb.RegisterBoookServiceServer(s, &server.Server{})

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
