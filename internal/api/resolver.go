package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/StevenACoffman/go-gql-todo/internal/model"
	"github.com/mitchellh/mapstructure"
	"time"
)

var you = &model.User{ID: 1, Name: "You"}
var them = &model.User{ID: 2, Name: "Them"}

func getUserId(ctx context.Context) int {
	if id, ok := ctx.Value("userId").(int); ok {
		return id
	}
	return you.ID
}

func New() Config {
	c := Config{
		Resolvers: &Resolver{
			todos: []*model.Todo{
				{ID: 1, Text: "A todo not to forget", Done: false, MyOwner: you},
				{ID: 2, Text: "This is the most important", Done: false, MyOwner: you},
				{ID: 3, Text: "Somebody else's todo", Done: true, MyOwner: them},
				{ID: 4, Text: "Please do this or else", Done: false, MyOwner: you},
			},
			lastID: 4,
		},
	}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		switch role {
		case model.RoleAdmin:
			// No admin for you!
			return nil, nil
		case model.RoleOwner:
			ownable, isOwnable := obj.(model.Ownable)
			if !isOwnable {
				return nil, fmt.Errorf("obj cant be owned")
			}

			if ownable.Owner().ID != getUserId(ctx) {
				return nil, fmt.Errorf("you dont own that")
			}
		}

		return next(ctx)
	}
	c.Directives.User = func(ctx context.Context, obj interface{}, next graphql.Resolver, id int) (interface{}, error) {
		return next(context.WithValue(ctx, "userId", id))
	}
	return c
}

type Resolver struct {
	todos  []*model.Todo
	lastID int
}

type QueryResolver Resolver

func (r *Resolver) MyQuery() MyQueryResolver {
	return (*QueryResolver)(r)
}

type MutationResolver Resolver

func (r *Resolver) MyMutation() MyMutationResolver {
	return (*MutationResolver)(r)
}

func (r *QueryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	time.Sleep(220 * time.Millisecond)

	if id == 666 {
		panic("critical failure")
	}

	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *QueryResolver) LastTodo(ctx context.Context) (*model.Todo, error) {
	if len(r.todos) == 0 {
		return nil, errors.New("not found")
	}
	return r.todos[len(r.todos)-1], nil
}

func (r *QueryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *MutationResolver) CreateTodo(ctx context.Context, todo model.TodoInput) (*model.Todo, error) {
	newID := r.id()

	newTodo := &model.Todo{
		ID:      newID,
		Text:    todo.Text,
		MyOwner: you,
	}

	if todo.Done != nil {
		newTodo.Done = *todo.Done
	}

	r.todos = append(r.todos, newTodo)

	return newTodo, nil
}

func (r *MutationResolver) UpdateTodo(ctx context.Context, id int, changes map[string]interface{}) (*model.Todo, error) {
	var affectedTodo *model.Todo

	for i := 0; i < len(r.todos); i++ {
		if r.todos[i].ID == id {
			affectedTodo = r.todos[i]
			break
		}
	}

	if affectedTodo == nil {
		return nil, nil
	}

	err := mapstructure.Decode(changes, affectedTodo)
	if err != nil {
		panic(err)
	}

	return affectedTodo, nil
}

func (r *MutationResolver) id() int {
	r.lastID++
	return r.lastID
}
