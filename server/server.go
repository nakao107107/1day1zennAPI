package server

import (
	"1day1zennAPI/middleware"
	"net/http"
	"github.com/gin-gonic/gin"
	history "1day1zennAPI/controller"
	user "1day1zennAPI/controller"
)

// 初期化
func Init() {
	r := router()
	r.Run()
}

// ルーティング
func router() *gin.Engine {
	r := gin.Default()

	// CORS対応
	r.Use(CORS())

	var m middleware.Middleware

	// ルーティング
	u := r.Group("")
	{
		historyController := history.Controller{}
		userController := user.Controller{}
		u.GET("/histories", m.CheckToken, historyController.Index)
		u.POST("/histories", m.CheckToken, historyController.Create)
		u.GET("/users/auth/login", userController.GetRedirectUrl)
		u.GET("/users/auth/callback", userController.GetToken)
	}

	return r
}

// CORS
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}