package main

import (
	"log"



	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"project/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
	"project/routes/auth"
)

const defaultPort = "8080"

func main() {

	app := fiber.New()

	
	auth.AuthApi(app)
	
	

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})


	app.All("/graphql",  adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	app.All("/query", adaptor.HTTPHandler(srv))
	
	
	
	log.Println("Server Works...")
	err := app.Listen(":8000")
	if err != nil {
		log.Print("Server Failed... main.go ln-15 \t", err)
	}
	
}
