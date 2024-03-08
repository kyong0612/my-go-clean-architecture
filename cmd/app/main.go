package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v5"
)

// const (
// 	defaultTimeout = 30
// 	defaultAddress = ":9090"
// )

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func main() {
	fmt.Println("Hello World")
	// //prepare database
	// dbHost := os.Getenv("DATABASE_HOST")
	// dbPort := os.Getenv("DATABASE_PORT")
	// dbUser := os.Getenv("DATABASE_USER")
	// dbPass := os.Getenv("DATABASE_PASS")
	// dbName := os.Getenv("DATABASE_NAME")
	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// val := url.Values{}
	// val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	// dbConn, err := sql.Open(`rdb`, dsn)
	// if err != nil {
	// 	log.Fatal("failed to open connection to database", err)
	// }
	// err = dbConn.Ping()
	// if err != nil {
	// 	log.Fatal("failed to ping database ", err)
	// }

	// defer func() {
	// 	err := dbConn.Close()
	// 	if err != nil {
	// 		log.Fatal("got error when closing the DB connection", err)
	// 	}
	// }()
	// // prepare echo

	// e := echo.New()
	// e.Use(middleware.CORS)
	// timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	// timeout, err := strconv.Atoi(timeoutStr)
	// if err != nil {
	// 	log.Println("failed to parse timeout, using default timeout")
	// 	timeout = defaultTimeout
	// }
	// timeoutContext := time.Duration(timeout) * time.Second
	// e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// // Prepare Repository
	// authorRepo := rdbRepo.NewAuthorRepository(dbConn)
	// articleRepo := rdbRepo.NewArticleRepository(dbConn)

	// // Build service Layer
	// svc := article.NewService(articleRepo, authorRepo)
	// rest.NewArticleHandler(e, svc)

	// // Start Server
	// address := os.Getenv("SERVER_ADDRESS")
	// if address == "" {
	// 	address = defaultAddress
	// }
	// log.Fatal(e.Start(os.Getenv("SERVER_ADDRESS"))) //nolint
}
