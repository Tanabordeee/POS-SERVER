package usecases

import "github.com/tanabordeee/pos/entity"

type ProductUseCase interface {
	CreateProduct(Product entity.Product) error
	GetProducts() ([]entity.Product, error)
	FindProductByName(name string) (entity.Product, error)
	UpdateUseCaseProduct(Product entity.Product) error
	DeleteUseCaseProduct(id uint) error
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductUseCase {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts() ([]entity.Product, error) {
	return s.repo.GetProduct()
}

func (s *ProductService) FindProductByName(name string) (entity.Product, error) {
	return s.repo.FindwithNameProduct(name)
}

func (s *ProductService) CreateProduct(product entity.Product) error {
	if err := s.repo.SaveProduct(product); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateUseCaseProduct(product entity.Product) error {
	if err := s.repo.UpdateProduct(product); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) DeleteUseCaseProduct(id uint) error {
	if err := s.repo.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}
