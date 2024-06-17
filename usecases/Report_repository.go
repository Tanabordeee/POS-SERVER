package usecases

import (
	"github.com/tanabordeee/pos/entity"
)

type ReportRepository interface {
	SaveReport(Report entity.Report) error
	GetReport(day int, month int, year int) ([]entity.Report, error)
	GetReport7Days() ([]entity.Report, error)
	GetReport1Month() ([]entity.Report, error)
	GetReport1Year() ([]entity.Report, error)
}
