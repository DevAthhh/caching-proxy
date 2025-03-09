package server

import (
	"github.com/DevAthhh/caching-proxy/app/cmd/server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Run(port, origin string) {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	router.GET("/*path", controllers.Index(origin))

	router.Run(":" + port)
}
