package routes

import (
	"final-project/controllers"
	"final-project/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoRouter(db *gorm.DB, router *gin.Engine) {
	var controller = controllers.InDB{
		DB: db,
	}

	photoRouterGroup := router.Group("/photos", middleware.Authentication(), middleware.JSONMiddleware())

	photoRouterGroup.POST("/", controller.Create)
	photoRouterGroup.GET("/", controller.Get)
	photoRouterGroup.PUT("/photoId", controller.UpdatePhoto)
	photoRouterGroup.DELETE("/photoId", controller.DeletePhoto)

}
