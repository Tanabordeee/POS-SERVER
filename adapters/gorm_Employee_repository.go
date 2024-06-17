package adapters

import (
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
	"gorm.io/gorm"
)

type GormEmployeeRepository struct {
	db *gorm.DB
}

func NewGormEmployeeRepository(db *gorm.DB) usecases.EmployeeRepository {
	return &GormEmployeeRepository{db: db}
}

func (r *GormEmployeeRepository) SaveEmployee(Employee entity.Employee) error {
	return r.db.Create(&Employee).Error
}

func (r *GormEmployeeRepository) FindEmployeeByName(Name string) (entity.Employee, error) {
	var Employee entity.Employee
	result := r.db.Preload("Auth").Where("name = ?", Name).First(&Employee)
	if result.Error != nil {
		return entity.Employee{}, result.Error
	}
	return Employee, nil
}

func (r *GormEmployeeRepository) FindEmployeeByID(id uint) error {
	var Employee entity.Employee
	result := r.db.First(&Employee, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormEmployeeRepository) GetEmployee() ([]entity.Employee, error) {
	var Employee []entity.Employee
	result := r.db.Preload("Auth").Find(&Employee)
	if result.Error != nil {
		return []entity.Employee{}, result.Error
	}
	return Employee, nil
}

func (r *GormEmployeeRepository) UpdateEmployee(employee *entity.Employee) error {
	result := r.db.Model(&entity.Employee{}).Where("employee_id = ?", employee.EmployeeID).Updates(employee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormEmployeeRepository) DeleteEmployee(id uint) error {
	var Employee entity.Employee
	result := r.db.Delete(&Employee, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
