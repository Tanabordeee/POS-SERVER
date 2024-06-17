package usecases

import (
	"github.com/tanabordeee/pos/entity"
)

type ReportUseCase interface {
	CreateReport(report []entity.Report) error
	GetReports(day int, month int, year int) ([]entity.Report, error)
	GetReports7Days() ([]entity.Report, error)
	GetReports1Month() ([]entity.Report, error)
	GetReports1Year() ([]entity.Report, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) ReportUseCase {
	return &ReportService{repo: repo}
}

func (s *ReportService) CreateReport(reports []entity.Report) error {
	for _, report := range reports {
		if err := s.repo.SaveReport(report); err != nil {
			return err
		}
	}
	return nil
}

func (s *ReportService) GetReports(day int, month int, year int) ([]entity.Report, error) {
	return s.repo.GetReport(day, month, year)
}

func (s *ReportService) GetReports7Days() ([]entity.Report, error) {
	return s.repo.GetReport7Days()
}

func (s *ReportService) GetReports1Month() ([]entity.Report, error) {
	return s.repo.GetReport1Month()
}

func (s *ReportService) GetReports1Year() ([]entity.Report, error) {
	return s.repo.GetReport1Year()
}
