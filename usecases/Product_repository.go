package usecases

import "github.com/tanabordeee/pos/entity"

type ProductRepository interface {
	SaveProduct(Product entity.Product) error
	FindwithNameProduct(Name string) (entity.Product, error)
	GetProduct() ([]entity.Product, error)
	UpdateProduct(Product entity.Product) error
	DeleteProduct(id uint) error
}
