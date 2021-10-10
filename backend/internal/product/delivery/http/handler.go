package http

import "github.com/ilyaGorkun/lineate-shop/internal/product"

type Handler struct {
	useCase product.Usecase
}
