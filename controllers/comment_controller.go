package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"final-project/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func (idb *InDB) CreateComment(ctx *gin.Context) {
	var commentReq utils.CommentRequest
	var comment models.Comment

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	errBindJson := ctx.ShouldBindJSON(&commentReq)

	if errBindJson != nil {
		ctx.JSON(400, gin.H{
			"result": nil,
			"err":    errBindJson,
			"data":   commentReq,
		})
		return
	}

	valid, trans := helpers.Validate()
	errValidate := valid.Struct(commentReq)

	if errValidate != nil {
		errs := errValidate.(validator.ValidationErrors)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Error Validation",
			"message": errs.Translate(trans),
		})
		return
	}

	comment = models.Comment{
		Message:   commentReq.Message,
		PhotoId:   commentReq.PhotoId,
		CreatedAt: time.Now(),
		UserId:    uint(userData["id"].(float64)),
	}

	err := idb.DB.Debug().Table("comments").Create(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to create comment",
			"err":     err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoId,
		"created_at": comment.CreatedAt,
	})
}

func (idb *InDB) GetComment(c *gin.Context) {
	var comments []utils.CommentResponse

	err := idb.DB.Debug().Table("comments").Preload("User").Preload("Photo").Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error Get Data",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (idb *InDB) UpdateComment(ctx *gin.Context) {
	var comment models.Comment
	var commentUpdateMsg utils.UpdateCommentMessage
	var commentId = ctx.Param("commentId")

	err := idb.DB.Table("comments").Where("id = ?", commentId).Take(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   `Can't get comment with id = ` + commentId,
			"message": "Error to update data!",
		})
		return
	}

	err = ctx.ShouldBindJSON(&commentUpdateMsg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"comments": nil,
			"err":      err,
		})
		return
	}

	valid, trans := helpers.Validate()
	errValidate := valid.Struct(commentUpdateMsg)

	if errValidate != nil {
		errs := errValidate.(validator.ValidationErrors)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Error Validation",
			"message": errs.Translate(trans),
		})
		return
	}

	idb.DB.Debug().Table("comments").Model(&comment).Where("id = ?", commentId).Updates(models.Comment{
		Message:   commentUpdateMsg.Message,
		UpdatedAt: time.Now(),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"photo_id":   comment.PhotoId,
		"message":    comment.Message,
		"user_id":    comment.UserId,
		"updated_at": comment.UpdatedAt,
	})

}

func (idb *InDB) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")

	errFind := idb.DB.Table("comments").Where("id = ?", commentId).First(models.Comment{}).Error

	if errFind != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error to delete comment",
			"message": "Unable to find comment with id = " + commentId,
		})
		return
	}

	err := idb.DB.Table("comments").Where("id = ?", commentId).Delete(models.Comment{}).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error to delete comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
