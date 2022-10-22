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
	routerGroupPhoto := router.Group("/photos", middleware.JSONMiddleware(), middleware.Authentication())
	routerGroupPhoto.GET("/", controller.GetPhoto)
	routerGroupPhoto.POST("/", controller.CreatePhoto)
	routerGroupPhoto.PUT("/:photoId", controller.UpdatePhoto)
	routerGroupPhoto.DELETE("/:photoId", controller.DeletePhoto)
}
