package usecases

import "github.com/tanabordeee/pos/entity"

type EmployeeRepository interface {
	SaveEmployee(Employee entity.Employee) error
	FindEmployeeByName(Name string) (entity.Employee, error)
	FindEmployeeByID(id uint) error
	GetEmployee() ([]entity.Employee, error)
	UpdateEmployee(Employee *entity.Employee) error
	DeleteEmployee(id uint) error
}
