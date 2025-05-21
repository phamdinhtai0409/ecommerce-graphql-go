package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ecommerce-graphql-go/data"
	"ecommerce-graphql-go/graph"
	"ecommerce-graphql-go/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize data
	data := data.NewData()

	// Generate mock tokens
	adminToken, _ := middleware.GenerateToken(data.GetUser("1"))
	customerToken, _ := middleware.GenerateToken(data.GetUser("2"))

	fmt.Println("Admin Token: ", adminToken)
	fmt.Println("Customer Token: ", customerToken)

	c := graph.Config{Resolvers: &graph.Resolver{Data: data}}
	c.Directives.HasRole = graph.HasRoleDirective

	srv := handler.New(graph.NewExecutableSchema(c))

	// Enable required transports
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	// DataLoader middleware
	serverLoaders := middleware.DataLoaderMiddleware(data, srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", middleware.AuthMiddleware(serverLoaders))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
