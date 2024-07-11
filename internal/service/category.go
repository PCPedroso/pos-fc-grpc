package service

import (
	"context"

	"github.com/PCPedroso/pos-fc-grpc/internal/database"
	"github.com/PCPedroso/pos-fc-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, r *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(r.Name, r.Description)
	if err != nil {
		return nil, err
	}

	// // Testar isto aqui depois
	// return &pb.CategoryResponse{
	// 	Category: &pb.Category{
	// 		Id:          category.ID,
	// 		Name:        category.Name,
	// 		Description: category.Description,
	// 	},
	// }, nil

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}
