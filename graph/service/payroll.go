package service

import (
	"backend-challenge/graph/model"
	"backend-challenge/graph/storage"
	"backend-challenge/util"
	"fmt"
	"github.com/golang/glog"
	"time"
)

type Payroll interface {
	GetPayroll(month, year int, country storage.Country) ([]*model.PayrollSummary, error)
}

type UserPayroll struct {
	store storage.Storage
}

func NewPayroll(store storage.Storage) *UserPayroll {
	return &UserPayroll{store: store}
}

func (payroll *UserPayroll) GetPayroll(month, year int, country storage.Country) ([]*model.PayrollSummary, error) {
	salaries, err := payroll.store.GetPayroll(month, year, country)
	if err != nil {
		return nil, err
	}
	taxes, err := payroll.store.GetTaxConfig(country)
	if err != nil {
		glog.Errorf("Couldn't load tax config for country %s %s", country, err)
	}
	var payrollSummary []*model.PayrollSummary
	for _, s := range salaries {
		payrollType := model.PayrollTypeReal
		payrollTime, _ := time.Parse("2006-01-02", fmt.Sprintf("%d-%d-%d", year, month, 1))
		if util.IsFutureDate(payrollTime) {
			payrollType = model.PayrollTypeFuturePreview
		}
		payrollSummary = append(payrollSummary, &model.PayrollSummary{
			Gross: s.Gross,
			Net:   s.Net,
			Bonus: &s.Bonus,
			User: &model.User{
				FirstName:         s.UserProfile.FirstName,
				LastName:          s.UserProfile.LastName,
				ProfilePictureURL: s.UserProfile.ProfilePictureUrl,
			},
			Taxes: getTaxes(taxes),
			Type:  payrollType,
		})
	}
	return payrollSummary, nil
}

func getTaxes(taxConfig []*storage.TaxConfig) []*model.Tax {
	var taxes []*model.Tax
	for _, t := range taxConfig {
		taxes = append(taxes, &model.Tax{
			Name:  t.TaxName,
			Value: t.TaxPercentageValue,
		})
	}
	return taxes
}
