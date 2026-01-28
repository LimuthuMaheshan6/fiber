package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"project/database"

	"project/graph"

	"project/routes/auth"
	"project/routes/contact"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	database.MongoConnection()
	/* defer database.CloseMongoConnection()
 */
	client := database.Client

	app := fiber.New()

	app.Use(cors.New())

	
	auth.AuthApi(app)
	contact.ContactRouter(app)
	
	
	

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Mongo: client,
		}}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})


	app.Get("/graphql",  adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	app.Post("/query", adaptor.HTTPHandler(srv))
	
	
	
	log.Println("Server Works...")
	err := app.Listen(":8000")
	if err != nil {
		log.Print("Server Failed... main.go ln-15 \t", err)
	}
	
}
