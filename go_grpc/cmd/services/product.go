package services

import (
	"context"
	pb "go_grpc/pb"

	"gorm.io/gorm"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductService) GetProducts(context.Context, *pb.Empty) (*pb.Products, error) {
	products := &pb.Products{
		Pagination: &pb.Pagination{
			Total:       10,
			PerPage:     5,
			CurrentPage: 1,
			LastPage:    2,
		},
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Metalica T-Shirt",
				Price: 100000.00,
				Stock: 10,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    1,
				Name:  "Black T-Shirt",
				Price: 300000.00,
				Stock: 10,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
		},
	}

	return products, nil
}
