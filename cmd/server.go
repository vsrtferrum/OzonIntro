package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vsrtferrum/OzonIntro/graph"

	"github.com/vsrtferrum/OzonIntro/internal/module"
	"github.com/vsrtferrum/OzonIntro/internal/readfromconfig"
	db "github.com/vsrtferrum/OzonIntro/internal/storage"
	"github.com/vsrtferrum/OzonIntro/internal/workers"
)

const defaultPort = "8080"
const pathToConfig = "config/config.json"
func main() {
    config, err := readfromconfig.ReadConfig(pathToConfig)
    if err != nil{
        panic(err)
    }

    var storage db.StorageAtions
    if config.DBStorage{
        storage, err  = db.NewDatabase(config.ConnStr)
        if err != nil{
            panic(err)
        }
    }else {
        storage = db.NewInMemoryStorage()
    }

	module := module.NewModule(storage)
    workersModule := workers.NewConcurrentModule(module, config.WorkersCount, config.WorkersQueueLen) 


    resolver := &graph.Resolver{
        Workers: workersModule,
    }


    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
        Resolvers: resolver,
    }))


    http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
    http.Handle("/query", srv)


    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    log.Printf("Server is running on http://localhost:%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}