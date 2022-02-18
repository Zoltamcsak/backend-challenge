package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend-challenge/graph/generated"
	"backend-challenge/graph/model"
	"backend-challenge/graph/storage"
	"context"
)

func (r *queryResolver) PayrollSummary(ctx context.Context, year int, month int, country model.Country) ([]*model.PayrollSummary, error) {
	s, _ := r.store.GetPayroll(month, year, storage.Country(country.String()))
	var res []*model.PayrollSummary
	for _, u := range s {
		res = append(res, &model.PayrollSummary{
			Gross: u.Gross,
			User: &model.User{
				FirstName: "",
			},
		})
	}
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
