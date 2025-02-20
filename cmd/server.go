package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vsrtferrum/OzonIntro/graph"

	"github.com/vsrtferrum/OzonIntro/internal/module"
	"github.com/vsrtferrum/OzonIntro/internal/storage"
	"github.com/vsrtferrum/OzonIntro/internal/workers"
)

const defaultPort = "8080"

func main() {
	storage :=  storage.NewInMemoryStorage()
	module := module.NewModule(storage)
    workersModule := workers.NewConcurrentModule(module, 10, 100) // Замените на вашу реальную инициализацию


    resolver := &graph.Resolver{
        Workers: workersModule,
    }

    // Создаем GraphQL-сервер
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
        Resolvers: resolver,
    }))

    // Настройка маршрутов
    http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
    http.Handle("/query", srv)

    // Запуск сервера
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    log.Printf("Server is running on http://localhost:%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}