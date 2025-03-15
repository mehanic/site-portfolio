package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", DBConnStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test database connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("Connected to PostgreSQL database!")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

//

// package config

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"

// 	_ "github.com/lib/pq"
// )

// var DB *sql.DB

// func InitDB() {
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")

// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		dbHost, dbPort, dbUser, dbPassword, dbName)

// 	var err error
// 	DB, err = sql.Open("postgres", dsn)
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to connect to DB: %v", err))
// 	}

// 	if err = DB.Ping(); err != nil {
// 		panic(fmt.Sprintf("DB is not reachable: %v", err))
// 	}

// 	fmt.Println("Connected to the database successfully")
// }

// func CloseDB() {
// 	if DB != nil {
// 		DB.Close()
// 	}
// }
