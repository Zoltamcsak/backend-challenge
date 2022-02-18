package main

import (
	"backend-challenge/graph"
	"backend-challenge/graph/generated"
	"backend-challenge/graph/service"
	"backend-challenge/graph/storage"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, dbName)),
		&gorm.Config{})
	if err != nil {
		glog.Fatal("Couldn't connect to database", err)
	}

	store := storage.NewDbStorage(db)
	db.AutoMigrate(&storage.UserProfile{}, &storage.Salary{}, &storage.TaxConfig{}, &storage.ExtraSalary{})

	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(service.NewPayroll(store))}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
