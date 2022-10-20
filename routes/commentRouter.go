package routes

import (
	"final-project/controllers"
	"final-project/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRouter(db *gorm.DB, router *gin.Engine) {
	var controller = controllers.InDB{
		DB: db,
	}
	routerGroupComment := router.Group("/comments", middleware.JSONMiddleware(), middleware.Authentication())
	routerGroupComment.GET("/", controller.GetComment)
	routerGroupComment.POST("/", controller.CreateComment)
	routerGroupComment.PUT("/:commentId", controller.UpdateComment)
	routerGroupComment.DELETE("/:commentId", controller.DeleteComment)
}
