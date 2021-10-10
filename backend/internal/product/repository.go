package product

import (
	"context"

	"github.com/ilyaGorkun/lineate-shop/models"
)

type Repository interface {
	GetProducts(сtx *context.Context) ([]*models.Product, error)
}
