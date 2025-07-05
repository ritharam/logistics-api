package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Failed to load .env")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil || DB.Ping() != nil {
		log.Fatal("❌ Failed to connect DB:", err)
	}
	log.Println("✅ DB connected")
	createTables()
}

func createTables() {
	shipmentTable := `CREATE TABLE IF NOT EXISTS shipments (
		id INT AUTO_INCREMENT PRIMARY KEY,
		origin VARCHAR(100),
		destination VARCHAR(100),
		weight FLOAT,
		urgency VARCHAR(50),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	if _, err := DB.Exec(shipmentTable); err != nil {
		log.Fatal("❌ Failed to create shipments table:", err)
	}
	log.Println("✅ Shipments table ensured")
}