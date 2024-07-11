package main

import (
	"database/sql"
	"net"

	"github.com/PCPedroso/pos-fc-grpc/internal/database"
	"github.com/PCPedroso/pos-fc-grpc/internal/pb"
	"github.com/PCPedroso/pos-fc-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./sb.sqlite")
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
