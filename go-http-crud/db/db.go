package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	var err error
	connectStr := os.Getenv("DB_STRING")
	DB, err = pgx.Connect(context.Background(), connectStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database successfully!")
}
