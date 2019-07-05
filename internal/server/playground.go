package server

import (
	"context"
	"github.com/StevenACoffman/go-gql-todo/internal/api"
	"log"
	"net/http"
	"runtime/debug"
	"errors"
	"github.com/99designs/gqlgen/handler"
)

// Handler returns a *http.ServeMux
func Handler() http.Handler {
	mux := http.NewServeMux()
	// Setup the playground
	mux.Handle("/", handler.Playground("Todo", "/query"))
	mux.Handle("/query", handler.GraphQL(
		api.NewExecutableSchema(api.New()),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			// send this panic somewhere
			log.Print(err)
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	))

	return mux
}
