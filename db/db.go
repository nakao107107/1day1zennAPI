package db

import (
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"1day1zennAPI/entity"
)

var (
	db *gorm.DB
	err error
)

// DB初期化
func Init() {
	// 環境変数取得
	godotenv.Load(".env")
	godotenv.Load()

	// DB接続
	db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

	if err != nil {
		panic(err)
	}

	autoMigration()
}

// DB取得
func GetDB() *gorm.DB {
	return db
}

// DB接続終了
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// マイグレーション
func autoMigration() {
	db.AutoMigrate(&entity.History{})
	db.AutoMigrate(&entity.User{})
}