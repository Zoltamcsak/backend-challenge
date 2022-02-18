package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend-challenge/graph/generated"
	"backend-challenge/graph/model"
	"backend-challenge/graph/storage"
	"context"
	"fmt"
)

func (r *mutationResolver) AddPayroll(ctx context.Context, userID int, country model.Country, grossSalary float64, year int, month int, bonus *float64) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PayrollSummary(ctx context.Context, year int, month int, country model.Country) ([]*model.PayrollSummary, error) {
	return r.payroll.GetPayroll(month, year, storage.Country(country.String()))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
