package main

import (
	"os"

	"example.com/go-api/request/validation"
	"example.com/go-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"example.com/go-api/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	godotenv.Load()

	validation.RegisterCustomValidation()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	// Set up GraphQL server
	gqlServer := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Add GraphQL transports
	gqlServer.AddTransport(transport.Options{})
	gqlServer.AddTransport(transport.GET{})
	gqlServer.AddTransport(transport.POST{})

	// Configure GraphQL query cache
	// gqlServer.SetQueryCache(lru.New(1000))

	// Enable GraphQL extensions
	gqlServer.Use(extension.Introspection{})
	gqlServer.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Mount GraphQL playground at / (optional, can be changed)
	server.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/graphql")))

	// Mount GraphQL query endpoint
	server.Any("/graphql", gin.WrapH(gqlServer))

	server.Run(":" + port)
}
