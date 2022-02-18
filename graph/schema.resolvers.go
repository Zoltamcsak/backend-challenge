package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend-challenge/graph/generated"
	"backend-challenge/graph/model"
	"backend-challenge/graph/storage"
	"context"
)

func (r *mutationResolver) AddPayroll(_ context.Context, data model.PayrollInput) (int, error) {
	return r.payrollService.SavePayroll(data)
}

func (r *queryResolver) PayrollSummary(ctx context.Context, year int, month int, country model.Country) ([]*model.PayrollSummary, error) {
	return r.payrollService.GetPayroll(month, year, storage.Country(country.String()))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
