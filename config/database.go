package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"tryg/models"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	
	_ = godotenv.Load()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Connect without DB to create it if missing
	rootDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)
	rawDB, err := sql.Open("mysql", rootDSN)
	if err != nil {
		return err
	}
	defer rawDB.Close()

	_, err = rawDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}
	log.Printf("Database %s ensured", dbName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Entry{}); err != nil {
		return err
	}

	DB = db
	log.Println("Database connection established and schema migrated")
	return nil
}