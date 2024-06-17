package adapters

import (
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) usecases.ProductRepository {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) SaveProduct(Product entity.Product) error {
	return r.db.Create(&Product).Error
}

func (r *GormProductRepository) GetProduct() ([]entity.Product, error) {
	var products []entity.Product
	result := r.db.Preload("Reports").Find(&products)
	if result.Error != nil {
		return []entity.Product{}, result.Error
	}
	return products, nil
}

func (r *GormProductRepository) FindwithNameProduct(Name string) (entity.Product, error) {
	var Product entity.Product
	result := r.db.Preload("Reports").Where("product_name = ?", Name).First(&Product)
	if result.Error != nil {
		return entity.Product{}, result.Error
	}
	return Product, nil
}

func (r *GormProductRepository) UpdateProduct(product entity.Product) error {
	result := r.db.Model(&entity.Product{}).Where("product_id = ?", product.ProductID).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormProductRepository) DeleteProduct(id uint) error {
	var Product entity.Product
	result := r.db.Delete(&Product, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
