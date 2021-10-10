package product

import (
	"context"

	"github.com/ilyaGorkun/lineate-shop/models"
)

type Usecase interface {
	GetProducts(сtx *context.Context) ([]*models.Product, error)
}
