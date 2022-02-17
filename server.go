package main

import (
	db2 "backend-challenge/db"
	"backend-challenge/graph"
	"backend-challenge/graph/generated"
	"backend-challenge/graph/storage"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	defaultPort = "8080"
	port        = os.Getenv("PORT")
	host        = os.Getenv("DB_HOST")
	user        = os.Getenv("DB_USER")
	pass        = os.Getenv("DB_PASS")
	dbName      = os.Getenv("DB_NAME")
)

func main() {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, dbName),
	)
	if err != nil {
		glog.Fatal("Couldn't connect to database", err)
	}
	db.SetMaxOpenConns(50)

	if err := db2.DoMigrations(db.DB); err != nil {
		glog.Fatal("Couldn't do db migration", err)
	}
	store := storage.NewDbStorage(db)

	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(store)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
