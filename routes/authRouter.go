package routes

import (
	"final-project/controllers"
	"final-project/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(db *gorm.DB, router *gin.Engine) {
	var controller = controllers.InDB{
		DB: db,
	}
	authRouterGroup := router.Group("/auth", middleware.JSONMiddleware())

	authRouterGroup.POST("/register", controller.Register)
	authRouterGroup.POST("/login", controller.Login)
	authRouterGroup.PUT("/update", middleware.Authentication(), controller.UpdateUser)
	authRouterGroup.DELETE("/delete", middleware.Authentication(), controller.DeleteUser)
}
