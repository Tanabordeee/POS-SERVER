package usecases

import "github.com/tanabordeee/pos/entity"

type DiscountRepository interface {
	SaveDiscount(Discount entity.Discount) error
	FindWithCode(code string) (entity.Discount, error)
	FindDiscountByID(id uint) error
	GetDiscount() ([]entity.Discount, error)
	UpdateDiscount(Discount entity.Discount) error
	DeleteDiscount(id uint) error
}
