package storage

import (
	"github.com/golang/glog"
	"gorm.io/gorm"
)

type Storage interface {
	GetPayroll(month, year int, country Country) ([]*Salary, error)
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

// DbStorage implements the Storage methods in memory as golang maps
type DbStorage struct {
	db *gorm.DB
}

// NewDbStorage returns a NewDbStorage with internal maps initialized
func NewDbStorage(db *gorm.DB) *DbStorage {
	return &DbStorage{db: db}
}
