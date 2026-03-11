package router

import (
	"github.com/110claw/backend/handler"
	"github.com/110claw/backend/middleware"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())

	r.GET("/110claw/api/health", handler.Health)

	v1 := r.Group("/110claw/api/v1")

	// Auth routes
	auth := v1.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/logout", middleware.Auth(), handler.Logout)
		auth.GET("/me", middleware.Auth(), handler.Me)
	}

	// Public content routes
	v1.GET("/news", handler.ListNews)
	v1.GET("/news/:id", handler.GetNews)

	v1.GET("/learn", handler.ListLearnSteps)
	v1.GET("/learn/:day", handler.GetLearnStep)

	v1.GET("/skills", handler.ListSkills)
	v1.GET("/skills/:id", handler.GetSkill)

	return r
}
