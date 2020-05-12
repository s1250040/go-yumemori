package server

import (
	"github.com/gin-gonic/gin"

	"github.com/s1250040/go-yumemori/controller/user_controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run() //サーバーを起動
}

func router() *gin.Engine {
	r := gin.Default() //ginの変数を定義

	u := r.Group("/users")
	{
		ctrl := user_controller.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
