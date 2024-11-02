package handlers

import (
	"context"
	"errors"

	"projgRPC/proto"
	"projgRPC/server/database"

	"gorm.io/gorm"
)

type ProductHandler struct {
	proto.UnimplementedProductServiceServer // Adicione esta linha
}

func (h *ProductHandler) AddProduct(ctx context.Context, req *proto.ProductRequest) (*proto.ProductResponse, error) {
	product := req.GetProduct()
	result := database.DB.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &proto.ProductResponse{Product: product}, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *proto.ProductID) (*proto.ProductResponse, error) {
	var product proto.Product
	result := database.DB.First(&product, req.GetId())
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return &proto.ProductResponse{Product: &product}, nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *proto.ProductRequest) (*proto.ProductResponse, error) {
	product := req.GetProduct()
	result := database.DB.Save(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &proto.ProductResponse{Product: product}, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *proto.ProductID) (*proto.DeleteResponse, error) {
	result := database.DB.Delete(&proto.Product{}, req.GetId())
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}
	return &proto.DeleteResponse{Message: "product deleted successfully"}, nil
}
