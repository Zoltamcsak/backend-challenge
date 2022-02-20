package storage

import (
	"github.com/golang/glog"
	"gorm.io/gorm"
)

type Storage interface {
	GetPayroll(month, year int, country Country) ([]*Salary, error)
	GetTaxConfig(country Country) ([]*TaxConfig, error)
	GetExtraSalary(country Country, month int) (*ExtraSalary, error)
	SaveSalary(salary *Salary) (uint, error)
}

func (d *DbStorage) GetPayroll(month, year int, country Country) ([]*Salary, error) {
	var salaries []*Salary
	tx := d.db.Preload("UserProfile").Where("month=? and year=? and country=?", month, year, country).Find(&salaries)
	if tx.Error != nil {
		glog.Error(tx.Error)
		return nil, tx.Error
	}
	return salaries, nil
}

func (d *DbStorage) GetTaxConfig(country Country) ([]*TaxConfig, error) {
	var taxConfig []*TaxConfig
	tx := d.db.Where("country=?", country).Find(&taxConfig)
	if tx.Error != nil {
		glog.Errorf("Error fetching tax config %s", tx.Error)
		return nil, tx.Error
	}
	return taxConfig, nil
}

func (d *DbStorage) SaveSalary(salary *Salary) (uint, error) {
	tx := d.db.Save(salary)
	if tx.Error != nil {
		return 0, tx.Error
	}
	result := tx.Statement.Model.(*Salary)
	return result.ID, nil
}

func (d *DbStorage) GetExtraSalary(country Country, month int) (*ExtraSalary, error) {
	var extra *ExtraSalary
	tx := d.db.Where("country=? and month=?", country, month).First(&extra)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return extra, nil
}

// DbStorage implements the Storage methods in memory as golang maps
type DbStorage struct {
	db *gorm.DB
}

// NewDbStorage returns a NewDbStorage with internal maps initialized
func NewDbStorage(db *gorm.DB) *DbStorage {
	return &DbStorage{db: db}
}
