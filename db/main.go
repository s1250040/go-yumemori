package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Users struct {
	Users []User `json:"users"`
}

func main() {
	db, err := gorm.Open("postgres", "user=yumemori password=yumemori　dbname=yumemori sslmode=disable")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You made it")
	}

	defer db.Close()

	r := gin.Default()
	// curl localhost:8080 とすると「Hello world」と表示する。
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	r.Run(":8080")

	// db.Set("gorm:table_options", "ENGINE=InnoDB")
	// db.AutoMigrate(&User{})
	db.LogMode(true)
}
