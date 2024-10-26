package main

import (
	"log"
	"net/http"

	"log/slog"

	graphql "github.com/graph-gophers/graphql-go"
	relay "github.com/graph-gophers/graphql-go/relay"
	_ "github.com/joho/godotenv/autoload"

	"gql-test/middleware"
	"gql-test/repository"
	"gql-test/resolvers"
)

func main() {
	config, err := GetConfig()
	slog.Info("main", "config", config)
	if err != nil {
		log.Fatal(err)
	}
	schema, err := config.GetSchema()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New()

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schemaHandler := graphql.MustParseSchema(schema, resolvers.New(repo), opts...)
	http.Handle("/graphql", middleware.LogRequest(&relay.Handler{Schema: schemaHandler}))
	slog.Info("started server")
	log.Fatal(http.ListenAndServe(config.Address, nil))
}
