package config

import (
	"database/sql"
	"fmt"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/pkg"
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabaseConnection() (db *sql.DB) {
	// Load ENV
	var (
		DB_HOST     = pkg.GetEnv("DB_HOST", "127.0.0.1")
		DB_PORT     = pkg.GetEnv("DB_PORT", "3306")
		DB_USERNAME = pkg.GetEnv("DB_USERNAME", "root")
		DB_PASSWORD = pkg.GetEnv("DB_PASSWORD", "")
		DB_NAME     = pkg.GetEnv("DB_NAME", "")
	)

	// Create Data Source Name
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	// Create Database Connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprint("Database Conenction Error, Err :", err.Error()))
	}

	// Ping Database Connection
	if err := db.Ping(); err != nil {
		db.Close()
		panic(fmt.Sprint("Database Connection Ping Error :", err.Error()))
	}

	return
}
