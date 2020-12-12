package persiterence

import (
	"zm-apd-go/domain/entity"
	"zm-apd-go/domain/repository"

	"github.com/jinzhu/gorm"
)

//ProductMapper mapper
type ProductMapper struct {
	db *gorm.DB
}

//NewProductMapper construct
func NewProductMapper(db *gorm.DB) *ProductMapper {
	return &ProductMapper{db}
}

var _ repository.ProductRepository = &ProductMapper{}

//CreateProduct insert one to db
func (pm *ProductMapper) CreateProduct(product *entity.Product) (*entity.Product, map[string]string) {
	dbErr := map[string]string{}
	error := pm.db.Debug().Create(&product).Error
	if error != nil {
		dbErr["db_error"] = "DB ERROR"
		return nil, dbErr
	}
	return product, nil
}
func (pm *ProductMapper) UpdateProduct(product *entity.Product) (*entity.Product, map[string]string) {
	dbErr := map[string]string{}
	findProduct, error := pm.QueryProduct(product.ID)
	if error != nil {
		dbErr["db_error"] = "DB ERROR"
		return nil, dbErr
	}
	if findProduct == nil {
		dbErr["db_error"] = "product not found"
		return nil, dbErr
	}
	err := pm.db.Debug().Table("products").Where("id=?", findProduct.ID).Update(&product).Error
	if err != nil {
		dbErr["db_error"] = "DB ERROR"
		return nil, dbErr
	}
	return product, nil

}
func (pm *ProductMapper) QueryProduct(id uint64) (*entity.Product, map[string]string) {
	var product entity.Product
	dbErr := map[string]string{}
	error := pm.db.Debug().Where("id=?", id).Take(&product).Error
	if error != nil {
		dbErr["db_error"] = "DB ERROR"
		return nil, dbErr
	}
	if gorm.IsRecordNotFoundError(error) {
		dbErr["db_error"] = "product not found"
		return nil, dbErr
	}
	return &product, nil
}
