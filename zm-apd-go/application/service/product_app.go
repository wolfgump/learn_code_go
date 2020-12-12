package service

import (
	"zm-apd-go/application/param"
	"zm-apd-go/domain/entity"
	"zm-apd-go/domain/repository"
)

//ProductApp product app
type ProductApp struct {
	repo repository.ProductRepository
}

//NewProductApp construct
func NewProductApp(repo repository.ProductRepository) *ProductApp {
	return &ProductApp{repo}
}

var _ ProductAppService = &ProductApp{}

//ProductAppService interface
type ProductAppService interface {
	CreateProduct(p *param.CreateProductCmd) bool
}

//CreateProduct impl
func (pa *ProductApp) CreateProduct(p *param.CreateProductCmd) bool {
	product := &entity.Product{}
	product.Name = p.Name
	product.Description = p.Description
	_, error := pa.repo.CreateProduct(product)
	if error != nil {
		return false
	}
	return true
}
func (pa *ProductApp) UpdateProduct(p *param.UpdateProductCmd) bool {
	product := &entity.Product{}
	product.ID = p.Id
	product.Name = p.Name
	product.Description = p.Description
	_, error := pa.repo.UpdateProduct(product)
	if error != nil {
		return false
	}
	return true
}
func (pa *ProductApp) Query(id uint64) *entity.Product {
	product, error := pa.repo.QueryProduct(id)
	if error != nil {
		return nil
	}
	return product
}
