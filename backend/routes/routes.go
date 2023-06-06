package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lits-06/go-auth/controllers"
)

func Setup(r *gin.Engine) {
	r.POST("/api/register", controllers.Register())
	r.POST("/api/login", controllers.Login())
	r.GET("/api/user", controllers.User())
	r.POST("/api/logout", controllers.Logout())
}
