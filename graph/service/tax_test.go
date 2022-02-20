package service

import (
	"backend-challenge/graph/storage"
	"testing"
)

func TestTaxService_GetNetSalary(t *testing.T) {
	type fields struct {
		store storage.Storage
	}
	type args struct {
		gross   float64
		country storage.Country
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:   "get net salary",
			fields: fields{store: &MockStorage{}},
			args:   args{country: storage.FRANCE, gross: 6000.0},
			want:   4242,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tax := &TaxService{
				store: tt.fields.store,
			}
			got, err := tax.GetNetSalary(tt.args.gross, tt.args.country)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNetSalary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNetSalary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxService_HasExtraSalary(t *testing.T) {
	type fields struct {
		store storage.Storage
	}
	type args struct {
		country storage.Country
		month   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "has extra salary",
			fields: fields{store: &MockStorage{}},
			args:   args{country: storage.ITALY, month: 12},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tax := &TaxService{
				store: tt.fields.store,
			}
			if got := tax.HasExtraSalary(tt.args.country, tt.args.month); got != tt.want {
				t.Errorf("HasExtraSalary() = %v, want %v", got, tt.want)
			}
		})
	}
}
