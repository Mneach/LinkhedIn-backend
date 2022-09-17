package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MneachDev/LinkhedIn-backend/authentication"
	"github.com/MneachDev/LinkhedIn-backend/graph"
	"github.com/MneachDev/LinkhedIn-backend/graph/generated"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// dsn := "host=localhost user=postgres password=mysecretpassword dbname=LinkhedIn port=12345 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=postgres password=admin dbname=LinkhedIn port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.ActivateAccount{},
		&model.ResetPasswordAccount{},
		&model.Education{},
		&model.Experience{},
		&model.Post{},
		&model.ConnectRequest{},
		&model.Connection{},
		&model.Comment{},
		&model.LikeComment{},
		&model.Hastag{},
		&model.Job{},
		&model.Notification{},
	)

	router := chi.NewRouter()

	router.Use(authentication.AuthMiddleware)
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		Debug:          true,
	}).Handler)

	c := generated.Config{Resolvers: &graph.Resolver{DB: db}}
	c.Directives.Auth = authentication.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "example.org"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
