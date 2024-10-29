package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"imersaofc/internal/converter"

	_ "github.com/lib/pq"
)

func connectPostgres() (*sql.DB, error) {
	user := getEnv0rDefault("POSTGRES_USER", "user")
	password := getEnv0rDefault("POSTGRES_PASSWORD", "password")
	dbname := getEnv0rDefault("POSTGRES_DB", "converter")
	host := getEnv0rDefault("POSTGRES_HOST", "postgres")
	sslmode := getEnv0rDefault("POSTGRES_SSL_MODE", "disable")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", user, password, dbname, host, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error("Error connecting to database", slog.String("connStr", connStr))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		slog.Error("Error pinging database", slog.String("connStr", connStr))
		return nil, err
	}
	slog.Info("Connected to Postgres successfuly")
	return db, nil
}

func getEnv0rDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	db, err := connectPostgres()
	if err != nil {
		panic(err)
	}

	vc := converter.NewVideoConverter(db)
	vc.Handle([]byte(`{"video_id": 1, "path": "/media/uploads/1"}`))
}
