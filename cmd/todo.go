package main

import (
	"context"
	"errors"
	"github.com/StevenACoffman/go-gql-todo/internal/api"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"
)

func main() {
	http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		api.NewExecutableSchema(api.New()),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			// send this panic somewhere
			log.Print(err)
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
