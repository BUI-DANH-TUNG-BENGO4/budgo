package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:password@tcp(localhost:3306)/app?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}

	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
}
