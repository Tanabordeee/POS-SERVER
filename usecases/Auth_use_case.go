package usecases

import (
	"github.com/tanabordeee/pos/entity"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	CheckAuth(Auth entity.Auth) (string, error)
	CreateAuth(Auth entity.Auth) error
	UpdateAuth(Auth entity.Auth) error
	DeleteAuth(id uint) error
}

type AuthService struct {
	AuthRepo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthUseCase {
	return &AuthService{AuthRepo: repo}
}

func (s *AuthService) CheckAuth(Auth entity.Auth) (string, error) {
	return s.AuthRepo.GetAuthRepository(Auth)
}

func (s *AuthService) CreateAuth(Auth entity.Auth) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	Auth.Password = string(hashedPassword)
	err = s.AuthRepo.SaveAuthRepository(Auth)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) UpdateAuth(Auth entity.Auth) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	Auth.Password = string(hashedPassword)
	err = s.AuthRepo.UpdateAuthRepository(Auth)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) DeleteAuth(id uint) error {
	if err := s.AuthRepo.DeleteAuthRepository(id); err != nil {
		return err
	}
	return nil
}
