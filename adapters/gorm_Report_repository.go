package adapters

import (
	"time"

	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
	"gorm.io/gorm"
)

type GormReportRepository struct {
	db *gorm.DB
}

func NewGormReportRepository(db *gorm.DB) usecases.ReportRepository {
	return &GormReportRepository{db: db}
}

func (r *GormReportRepository) SaveReport(Report entity.Report) error {
	return r.db.Create(&Report).Error
}

func (r *GormReportRepository) GetReport(day int, month int, year int) ([]entity.Report, error) {
	var reports []entity.Report

	// สร้างตัวแปรวันที่เริ่มต้นและสิ้นสุดของช่วงเวลา
	startDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 1).Add(-time.Microsecond)

	// สร้าง query ด้วยช่วงเวลาที่กำหนดและ preload ข้อมูล Product
	result := r.db.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("product_id, product_name, price, cost, created_at, updated_at, deleted_at")
	}).Where("report_date BETWEEN ? AND ?", startDate, endDate).Find(&reports)
	if result.Error != nil {
		return nil, result.Error
	}

	return reports, nil
}

func (r *GormReportRepository) GetReport7Days() ([]entity.Report, error) {
	var reports []entity.Report

	// สร้างตัวแปรวันที่เริ่มต้นและสิ้นสุดของช่วงเวลา (7 วันย้อนหลัง)
	endDate := time.Now().UTC()
	startDate := endDate.AddDate(0, 0, -7)

	// สร้าง query ด้วยช่วงเวลาที่กำหนด
	result := r.db.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("product_id, product_name, price, cost, created_at, updated_at, deleted_at")
	}).Where("report_date BETWEEN ? AND ?", startDate, endDate).Find(&reports)
	if result.Error != nil {
		return nil, result.Error
	}

	return reports, nil
}

func (r *GormReportRepository) GetReport1Month() ([]entity.Report, error) {
	var reports []entity.Report

	// สร้างตัวแปรวันที่เริ่มต้นและสิ้นสุดของช่วงเวลา (1 เดือนย้อนหลัง)
	endDate := time.Now().UTC()
	startDate := endDate.AddDate(0, -1, 0)

	// สร้าง query ด้วยช่วงเวลาที่กำหนด
	result := r.db.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("product_id, product_name, price, cost, created_at, updated_at, deleted_at")
	}).Where("report_date BETWEEN ? AND ?", startDate, endDate).Find(&reports)
	if result.Error != nil {
		return nil, result.Error
	}

	return reports, nil
}

func (r *GormReportRepository) GetReport1Year() ([]entity.Report, error) {
	var reports []entity.Report

	// สร้างตัวแปรวันที่เริ่มต้นและสิ้นสุดของช่วงเวลา (1 ปีย้อนหลัง)
	endDate := time.Now().UTC()
	startDate := endDate.AddDate(-1, 0, 0)

	// สร้าง query ด้วยช่วงเวลาที่กำหนด
	result := r.db.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("product_id, product_name, price, cost, created_at, updated_at, deleted_at")
	}).Where("report_date BETWEEN ? AND ?", startDate, endDate).Find(&reports)
	if result.Error != nil {
		return nil, result.Error
	}

	return reports, nil
}
