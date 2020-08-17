package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/s1250040/go-yumemori/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	// db, err = gorm.Open("postgres", "user=yumemori password=yumemori　dbname=yumemori sslmode=disable") test時
	// db, err = gorm.Open("postgres", "host=163.143.158.118 port=5432 user=postgres dbname=usa-test password=12345 sslmode=disable") //開発用
	db, err = gorm.Open("postgres", "host=163.143.158.114 port=5551 user=postgres dbname=usa-test sslmodel=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	// db.AutoMigrate(&entity.User{}, &entity.State{})　test時
	db.AutoMigrate(&entity.User{})
}
