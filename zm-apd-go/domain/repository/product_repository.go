package repository

import (
	"zm-apd-go/domain/entity"
)

//ProductRepository product repository
type ProductRepository interface {
	CreateProduct(p *entity.Product) (*entity.Product, map[string]string)
	UpdateProduct(p *entity.Product) (*entity.Product, map[string]string)
	QueryProduct(id uint64) (*entity.Product, map[string]string)
}
