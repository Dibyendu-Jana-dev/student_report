package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"mywon/students_reports/graph/model"
	"mywon/students_reports/logic"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return nil, nil
}

// UpsertStudentDetails is the resolver for the upsertStudentDetails field.
func (r *mutationResolver) UpsertStudentDetails(ctx context.Context, input model.CreateStudentsInput) (*model.CreateStudentsResponse, error) {
	resp, err := logic.CreateStudentDetails(ctx, input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//panic(fmt.Errorf("not implemented: Todos - todos"))
	return nil, nil
}

// GetStudentDetails is the resolver for the getStudentDetails field.
func (r *queryResolver) GetStudentDetails(ctx context.Context, input model.GetStudentDetailsInput) (*model.CreateStudentsResponse, error) {
	resp, err := logic.GetStudentDetails(ctx, input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
