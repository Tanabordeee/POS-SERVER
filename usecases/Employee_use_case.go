package usecases

import (
	"errors"

	"github.com/tanabordeee/pos/entity"
)

type EmployeeUseCase interface {
	CreateEmployee(Employee entity.Employee) error
	FindEmployee(Name string) (entity.Employee, error)
	GetEmployees() ([]entity.Employee, error)
	UpdateUseCaseEmployee(Employee entity.Employee) error
	DeleteUseCaseEmployee(id uint) error
}

type EmployeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) EmployeeUseCase {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee entity.Employee) error {
	if employee.Name == "" {
		return errors.New("EMPLOYEE NAME CANNOT BE EMPTY")
	}
	if employee.Role == "" {
		return errors.New("EMPLOYEE ROLE CANNOT BE EMPTY")
	}
	if employee.Salary <= 0 {
		return errors.New("EMPLOYEE SALARY MUST BE GREATER THAN ZERO")
	}
	if err := s.repo.SaveEmployee(employee); err != nil {
		return err
	}
	return nil
}

func (s *EmployeeService) FindEmployee(name string) (entity.Employee, error) {
	if name == "" {
		return entity.Employee{}, errors.New("NAME CANNOT FOUND")
	}
	return s.repo.FindEmployeeByName(name)
}

func (s *EmployeeService) GetEmployees() ([]entity.Employee, error) {
	return s.repo.GetEmployee()
}

func (s *EmployeeService) UpdateUseCaseEmployee(Employee entity.Employee) error {
	if Employee.Name == "" {
		return errors.New("EMPLOYEE NAME CANNOT BE EMPTY")
	}
	if Employee.Role == "" {
		return errors.New("EMPLOYEE ROLE CANNOT BE EMPTY")
	}
	if Employee.Salary <= 0 {
		return errors.New("EMPLOYEE SALARY MUST BE GREATER THAN ZERO")
	}
	if err := s.repo.UpdateEmployee(&Employee); err != nil {
		return err
	}
	return nil
}

func (s *EmployeeService) DeleteUseCaseEmployee(id uint) error {
	if err := s.repo.FindEmployeeByID(id); err != nil {
		return errors.New("WE DON'T HAVE THAT EMPLOYEE")
	}
	if err := s.repo.DeleteEmployee(id); err != nil {
		return err
	}
	return nil
}
