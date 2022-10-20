package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	AuthRouter(db, router)
	PhotoRouter(db, router)
	CommentRouter(db, router)
	SocialMediaRouter(db, router)

	return router
}
