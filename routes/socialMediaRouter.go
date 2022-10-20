package routes

import (
	"final-project/controllers"
	"final-project/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SocialMediaRouter(db *gorm.DB, router *gin.Engine) {
	var controller = controllers.InDB{
		DB: db,
	}
	socialMediaRouter := router.Group("/socialmedias", middleware.JSONMiddleware(), middleware.Authentication())
	socialMediaRouter.GET("/", controller.GetSocialMedia)
	socialMediaRouter.POST("/", controller.CreateSocialMedia)
	socialMediaRouter.PUT("/:socialMediaId", controller.UpdateSocialMedia)
	socialMediaRouter.DELETE("/:socialMediaId", controller.DeleteSocialMedia)
}
