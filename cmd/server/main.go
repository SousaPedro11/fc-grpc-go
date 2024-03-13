package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/sousapedro11/fc-grpc-go/internal/database"
	"github.com/sousapedro11/fc-grpc-go/internal/pb"
	"github.com/sousapedro11/fc-grpc-go/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}
