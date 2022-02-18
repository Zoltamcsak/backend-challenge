package graph

//go:generate go run github.com/99designs/gqlgen

import "backend-challenge/graph/storage"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store storage.Storage
}

// NewResolver returns a Resolver
func NewResolver(store storage.Storage) *Resolver {
	output := &Resolver{store: store}
	return output
}
