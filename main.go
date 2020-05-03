package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// main
func main() {
	r := gin.Default()
	// curl localhost:8080 とすると「Hello world」と表示する。
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	r.Run(":8080")
}
