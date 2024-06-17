package usecases

import "github.com/tanabordeee/pos/entity"

type AuthRepository interface {
	GetAuthRepository(Auth entity.Auth) (string, error)
	SaveAuthRepository(Auth entity.Auth) error
	UpdateAuthRepository(Auth entity.Auth) error
	DeleteAuthRepository(id uint) error
}
