package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseHandler struct {
  DB *gorm.DB
}

func NewDatabaseHandler() *DatabaseHandler {
	 dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database",err)
	}
	if os.Getenv("ENV") == "development" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} 

	return &DatabaseHandler{DB: db}
}

func (h *DatabaseHandler) Close() {
	sqlDB, _ := h.DB.DB()
	sqlDB.Close()
}

func (h *DatabaseHandler) Conn(ctx context.Context) *gorm.DB {
	return h.DB.WithContext(ctx)
}