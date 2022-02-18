package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"backend-challenge/graph/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	payrollService service.Payroll
}

// NewResolver returns a Resolver
func NewResolver(payroll service.Payroll) *Resolver {
	output := &Resolver{payrollService: payroll}
	return output
}
