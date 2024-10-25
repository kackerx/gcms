package router

import (
	"github.com/gin-gonic/gin"

	"gcms/internal/handler"
)

func RegisterUserRoute(r *gin.Engine, userHandler *handler.UserHandler) {
	user := r.Group("/user")
	user.GET("/hello", userHandler.GetUser)
}
