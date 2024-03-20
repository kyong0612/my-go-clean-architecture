package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kyong0612/my-go-clean-architecture/internal/repository"
	"github.com/kyong0612/my-go-clean-architecture/internal/rest"
	"github.com/kyong0612/my-go-clean-architecture/internal/rest/middleware"
	"github.com/kyong0612/my-go-clean-architecture/usecase"
	"github.com/labstack/echo/v4"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s database=%s sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal("failed to parse config", err)
	}

	dbConn, err := pgx.ConnectConfig(context.TODO(), config)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	defer dbConn.Close(context.TODO())

	// prepare echo

	e := echo.New()
	e.Use(middleware.CORS)

	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")

	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")

		timeout = defaultTimeout
	}

	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// Prepare Repository
	repo := repository.New(dbConn)

	// Build service Layer
	svc := usecase.NewService(repo)
	rest.NewArticleHandler(e, svc)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}

	log.Fatal(e.Start(address)) //nolint
}
