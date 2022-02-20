package service

import (
	"backend-challenge/graph/model"
	"backend-challenge/graph/storage"
	"reflect"
	"testing"
)

type MockStorage struct {
	savedSalary *storage.Salary
}

type MockTaxService struct{}

func (ms *MockStorage) GetPayroll(_, _ int, _ storage.Country) ([]*storage.Salary, error) {
	var salaries []*storage.Salary
	salaries = append(salaries, &storage.Salary{
		Country: storage.FRANCE,
		Gross:   5000,
		Bonus:   500,
		Net:     3800,
		ID:      1,
		UserID:  74,
		UserProfile: storage.UserProfile{
			FirstName:         "T",
			LastName:          "Z",
			ProfilePictureUrl: "www.profile.com",
		},
	})
	return salaries, nil
}

func (ms *MockStorage) GetTaxConfig(country storage.Country) ([]*storage.TaxConfig, error) {
	var taxConfig []*storage.TaxConfig
	taxConfig = append(taxConfig, &storage.TaxConfig{
		ID:                 1,
		TaxPercentageValue: 29.3,
		TaxName:            "tax",
		Country:            country,
	})
	return taxConfig, nil
}

func (ms *MockStorage) GetExtraSalary(c storage.Country, _ int) (*storage.ExtraSalary, error) {
	if c == storage.ITALY {
		return &storage.ExtraSalary{}, nil
	}
	return nil, nil
}

func (ms *MockStorage) SaveSalary(s *storage.Salary) (uint, error) {
	ms.savedSalary = s
	return 0, nil
}

func (t *MockTaxService) GetNetSalary(_ float64, _ storage.Country) (float64, error) {
	return 5000.0, nil
}
func (t *MockTaxService) HasExtraSalary(_ storage.Country, _ int) bool {
	return true
}

func TestPayrollService_GetPayroll(t *testing.T) {
	var bonus *float64
	bonus = new(float64)
	*bonus = 500
	type fields struct {
		store      storage.Storage
		taxService Tax
	}
	type args struct {
		month   int
		year    int
		country storage.Country
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.PayrollSummary
		wantErr bool
	}{
		{
			name:   "get payroll summary",
			fields: fields{store: &MockStorage{}, taxService: &MockTaxService{}},
			args:   args{month: 12, year: 2021, country: storage.FRANCE},
			want: []*model.PayrollSummary{{
				Net:   3800,
				Type:  model.PayrollTypeReal,
				Gross: 5000,
				Bonus: bonus,
				User: &model.User{
					FirstName:         "T",
					LastName:          "Z",
					ProfilePictureURL: "www.profile.com",
				},
				Taxes: []*model.Tax{{Name: "tax", Value: 29.3}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payroll := &PayrollService{
				store:      tt.fields.store,
				taxService: tt.fields.taxService,
			}
			got, err := payroll.GetPayroll(tt.args.month, tt.args.year, tt.args.country)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPayroll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPayroll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayrollService_SavePayroll(t *testing.T) {
	var bonus *float64
	bonus = new(float64)
	*bonus = 0.0
	t.Run("save payroll", func(t *testing.T) {
		mockTax := &MockTaxService{}
		mockStore := &MockStorage{}
		payroll := &PayrollService{
			store:      mockStore,
			taxService: mockTax,
		}

		payroll.SavePayroll(model.PayrollInput{
			UserID: 4, Year: 2022, Month: 3, Bonus: bonus, Country: model.CountryFrance, GrossSalary: 6000.0,
		})
		expectedSalary := &storage.Salary{
			Country: storage.FRANCE,
			Net:     5000.0,
			Gross:   12000.0,
			Bonus:   0.0,
			Month:   3,
			Year:    2022,
			UserID:  4,
		}
		if expectedSalary.Net != mockStore.savedSalary.Net {
			t.Errorf("SavePayroll() got = %v, want %v", mockStore.savedSalary.Net, expectedSalary.Net)
		}
		if expectedSalary.Gross != mockStore.savedSalary.Gross {
			t.Errorf("SavePayroll() got = %v, want %v", mockStore.savedSalary.Gross, expectedSalary.Gross)
		}
	})
}
