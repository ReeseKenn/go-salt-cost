package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "pg-for-go-salt-cost"
	port     = 5432
	user     = "salt"
	password = "123456"
	dbname   = "salt_cost_go_dev"
)

func Connect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connect to db")
}

func CreateTables() {
	// 创建 users 表
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully create users table")
}

func Close() {
	DB.Close()
	log.Println("Successfully close db")
}
