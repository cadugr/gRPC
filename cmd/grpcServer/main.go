package main

import (
	"database/sql"
	"net"

	"github.com/cadugr/gRPC/internal/database"
	"github.com/cadugr/gRPC/internal/pb"
	"github.com/cadugr/gRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", ".db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	//Registrando o nosso service no nosso servidor gRPC:
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer) // Permitir usarmos um client como o Evans.
	// Agora, precisamos abrir uma conexão TCP para "falar" com o gRPC:
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
