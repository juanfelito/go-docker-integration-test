package main

import (
	"context"
	"docker-example/internal/database/postgres"
	"docker-example/pkg/controllers"
	"docker-example/pkg/mediators"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
	"strconv"
	"strings"

	"docker-example/pkg/config"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var settings config.Settings
	if err := envconfig.Process("", &settings); err != nil {
		panic(fmt.Errorf("parsing environment variables into settings struct: %w", err))
	}

	dbProperties := []string{
		"host=" + settings.DBHost,
		"port=" + strconv.Itoa(settings.DBPort),
		"user=" + settings.DBUser,
		"dbname=" + settings.DBName,
		"sslmode=" + settings.DBSSLMode,
	}
	if settings.DBPass != "" {
		dbProperties = append(dbProperties, "password="+settings.DBPass)
	}

	dbConfig, err := pgxpool.ParseConfig(strings.Join(dbProperties, " "))
	if err != nil {
		panic(fmt.Errorf("failed to parse database configuration: %w", err))
	}

	dbPool, err := pgxpool.ConnectConfig(context.Background(), dbConfig)
	if err != nil {
		panic(fmt.Errorf("failed to initialize database pool: %w", err))
	}

	db := postgres.NewDB(dbPool)

	messageMediator := mediators.NewMessageMediator(mediators.WithDatabase(db))
	messagesController := controllers.NewMessagesController(messageMediator)

	r := mux.NewRouter()
	r.Path("/message/{id}").Methods(http.MethodGet).HandlerFunc(messagesController.GetMessage)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", settings.HTTPPort),
		Handler: r,
	}

	fmt.Println("Server running on port ", settings.HTTPPort)
	srv.ListenAndServe()
}
