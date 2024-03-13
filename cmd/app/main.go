package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//prepare database
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	connURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-full", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := pgx.Connect(context.TODO(), connURL)
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
	log.Fatal(e.Start(os.Getenv("SERVER_ADDRESS"))) //nolint
}
