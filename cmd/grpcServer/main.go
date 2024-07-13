package main

import (
	"database/sql"
	"net"

	"github.com/PCPedroso/pos-fc-grpc/internal/database"
	"github.com/PCPedroso/pos-fc-grpc/internal/pb"
	"github.com/PCPedroso/pos-fc-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

// Criar o banco de dados sqlite3, conectar e em seguida criar o banco
// sqlite3 data.db
// CREATE TABLE categories (id string, name string, description string);
//
// Evans install
// go install github.com/ktr0731/evans@latest
// Com a aplicação rodando executar em um novo terminal a instrução, e em seguida fazer uma call para criar um registro
// evans -r repl
// Para utilizar um service primeiro exetuar
// package pb
// Em seguida setar o service como no ex
// service CategoryService
// call CreateCategory

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(list); err != nil {
		panic(err)
	}
}
