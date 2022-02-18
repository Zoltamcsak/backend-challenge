package service

import (
	"backend-challenge/graph/storage"
	"github.com/golang/glog"
)

type Tax interface {
	GetNetSalary(gross float64, country storage.Country) (float64, error)
	HasExtraSalary(country storage.Country, month int) bool
}

type TaxService struct {
	store storage.Storage
}

func NewTaxService(store storage.Storage) *TaxService {
	return &TaxService{store: store}
}

func (tax *TaxService) GetNetSalary(gross float64, country storage.Country) (float64, error) {
	taxConfig, err := tax.store.GetTaxConfig(country)
	if err != nil {
		glog.Errorf("Couldn't get tax config for country %s %s", country, err)
		return 0, err
	}
	totalTax := 0.0
	for _, t := range taxConfig {
		totalTax += t.TaxPercentageValue
	}
	if totalTax != 0.0 {
		return gross - (gross * totalTax / 100), nil
	}
	return gross, nil
}

func (tax *TaxService) HasExtraSalary(country storage.Country, month int) bool {
	extraSalary, _ := tax.store.GetExtraSalary(country, month)
	return extraSalary != nil
}
