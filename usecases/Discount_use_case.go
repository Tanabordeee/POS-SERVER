package usecases

import (
	"errors"

	"github.com/tanabordeee/pos/entity"
)

type DiscountUseCase interface {
	CreateDiscount(Discount entity.Discount) error
	FindDiscountByCode(Code string) (entity.Discount, error)
	GetDiscounts() ([]entity.Discount, error)
	UpdateUseCaseDiscount(Discount entity.Discount) error
	DeleteUseCaseDiscount(id uint) error
}

type DiscountService struct {
	repo DiscountRepository
}

func NewDiscountService(repo DiscountRepository) DiscountUseCase {
	return &DiscountService{repo: repo}
}

func (s *DiscountService) CreateDiscount(Discount entity.Discount) error {
	if Discount.DiscountCode == "" {
		return errors.New("DISCOUNT CODE CANNOT BE EMPTY")
	}
	if err := s.repo.SaveDiscount(Discount); err != nil {
		return err
	}
	return nil
}

func (s *DiscountService) FindDiscountByCode(Code string) (entity.Discount, error) {
	if Code == "" {
		return entity.Discount{}, errors.New("CODE CANNOT FOUND")
	}
	return s.repo.FindWithCode(Code)
}

func (s *DiscountService) GetDiscounts() ([]entity.Discount, error) {
	return s.repo.GetDiscount()
}

func (s *DiscountService) UpdateUseCaseDiscount(Discount entity.Discount) error {
	if Discount.DiscountCode == "" {
		return errors.New("DISCOUNT CODE CANNOT BE EMPTY")
	}
	if err := s.repo.UpdateDiscount(Discount); err != nil {
		return err
	}
	return nil
}

func (s *DiscountService) DeleteUseCaseDiscount(id uint) error {
	if err := s.repo.FindDiscountByID(id); err != nil {
		return errors.New("WE DON'T HAVE THAT DISCOUNT")
	}
	if err := s.repo.DeleteDiscount(id); err != nil {
		return err
	}
	return nil
}
