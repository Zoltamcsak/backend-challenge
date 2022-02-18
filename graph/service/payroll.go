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
	SavePayroll(data model.PayrollInput) (int, error)
}

type PayrollService struct {
	store      storage.Storage
	taxService Tax
}

func NewPayroll(store storage.Storage, taxService Tax) *PayrollService {
	return &PayrollService{store: store, taxService: taxService}
}

func (payroll *PayrollService) GetPayroll(month, year int, country storage.Country) ([]*model.PayrollSummary, error) {
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

func (payroll *PayrollService) SavePayroll(data model.PayrollInput) (int, error) {
	country := storage.GetCountry(data.Country)
	extraSalary := payroll.taxService.HasExtraSalary(country, data.Month)
	if extraSalary {
		data.GrossSalary = data.GrossSalary * 2
	}

	netSalary, err := payroll.taxService.GetNetSalary(data.GrossSalary, country)
	if err != nil {
		netSalary = data.GrossSalary
	}

	bonus := 0.0
	if data.Bonus != nil {
		bonus = *data.Bonus
	}

	salary := &storage.Salary{
		Country: country,
		Net:     netSalary,
		Gross:   data.GrossSalary,
		Bonus:   bonus,
		Month:   data.Month,
		Year:    data.Year,
		UserID:  uint(data.UserID),
	}
	if err := payroll.store.SaveSalary(salary); err != nil {
		glog.Errorf("Couldn't save payroll %s", err)
		return 0, err
	}
	return 0, nil
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
