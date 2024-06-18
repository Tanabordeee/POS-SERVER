package adapters

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) usecases.AuthRepository {
	return &GormAuthRepository{db: db}
}

func (r *GormAuthRepository) GetAuthRepository(Auth entity.Auth) (string, error) {
	var SelectAuth entity.Auth
	result := r.db.Where("username = ?", Auth.Username).First(&SelectAuth)

	if result.Error != nil {
		return "", result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(SelectAuth.Password), []byte(Auth.Password))
	if err != nil {
		return "", err
	}
	// pass = return jwt
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["auth_id"] = SelectAuth.AuthID
	claims["auth_name"] = SelectAuth.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWTSECRETKEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (r *GormAuthRepository) SaveAuthRepository(Auth entity.Auth) error {
	return r.db.Create(&Auth).Error
}

func (r *GormAuthRepository) UpdateAuthRepository(Auth entity.Auth) error {
	result := r.db.Model(&entity.Auth{}).Where("auth_id = ?", Auth.AuthID).Updates(Auth)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormAuthRepository) DeleteAuthRepository(id uint) error {
	var Auth entity.Auth
	result := r.db.Delete(&Auth, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
