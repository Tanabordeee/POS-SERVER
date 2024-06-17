package adapters

import (
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
	"gorm.io/gorm"
)

type GormDiscountRepository struct {
	db *gorm.DB
}

func NewGormDiscountRepository(db *gorm.DB) usecases.DiscountRepository {
	return &GormDiscountRepository{db: db}
}

func (r *GormDiscountRepository) SaveDiscount(Discount entity.Discount) error {
	return r.db.Create(&Discount).Error
}

func (r *GormDiscountRepository) FindWithCode(code string) (entity.Discount, error) {
	var Discount entity.Discount
	result := r.db.Where("discount_code = ?", code).First(&Discount)
	if result.Error != nil {
		return entity.Discount{}, result.Error
	}
	return Discount, nil
}

func (r *GormDiscountRepository) FindDiscountByID(id uint) error {
	var Discount entity.Discount
	result := r.db.Where("discount_id = ?", id).First(&Discount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormDiscountRepository) GetDiscount() ([]entity.Discount, error) {
	var Discount []entity.Discount
	result := r.db.Find(&Discount)
	if result.Error != nil {
		return []entity.Discount{}, result.Error
	}
	return Discount, nil
}

func (r *GormDiscountRepository) UpdateDiscount(Discount entity.Discount) error {
	result := r.db.Model(&entity.Discount{}).Where("discount_id = ?", Discount.DiscountID).Updates(Discount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormDiscountRepository) DeleteDiscount(id uint) error {
	var Discount entity.Discount
	result := r.db.Delete(&Discount, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
