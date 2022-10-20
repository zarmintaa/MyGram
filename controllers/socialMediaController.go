package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type SocialMediaRequest struct {
	Name             string `json:"name" validate:"required,max=50"`
	Social_media_url string `json:"social_media_url" validate:"required,max=191"`
}

type SocialMediaResponse struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Social_media_url string `json:"social_media_url"`
	Created_at       string `json:"created_at"`
	Updated_at       string `json:"updated_at"`
	User_id          string `json:"-"`
	User             *User  `json:"user"`
}

func (idb *InDB) CreateSocialMedia(ctx *gin.Context) {
	var SocialMedia models.SocialMedia
	var SocialMediaReq SocialMediaRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	err := ctx.ShouldBindJSON(&SocialMediaReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"err":    err,
		})
		return
	}

	valid, trans := helpers.Validate()
	err = valid.Struct(SocialMediaReq)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Error Validation",
			"message": errs.Translate(trans),
		})
		return
	}

	SocialMedia = models.SocialMedia{
		Name:             SocialMediaReq.Name,
		Social_media_url: SocialMediaReq.Social_media_url,
		Created_at:       time.Now(),
		User_id:          uint(userData["id"].(float64)),
	}

	err = idb.DB.Debug().Table("social_media").Create(&SocialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error to create comment",
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.Id,
		"message":          SocialMedia.Name,
		"social_media_url": SocialMedia.Social_media_url,
		"user_id":          SocialMedia.User_id,
		"created_at":       SocialMedia.Created_at,
	})
}

func (idb *InDB) GetSocialMedia(ctx *gin.Context) {
	var listSocialMedia []SocialMediaResponse

	err := idb.DB.Debug().Table("social_media").Preload("User").Find(&listSocialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error Get Social",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, listSocialMedia)
}

func (idb *InDB) UpdateSocialMedia(ctx *gin.Context) {
	var socialMediaRequest SocialMediaRequest
	var socialMedia models.SocialMedia
	socialMediaId := ctx.Param("socialMediaId")

	errJson := ctx.ShouldBindJSON(&socialMediaRequest)

	if errJson != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"err":    errJson,
		})
		return
	}

	valid, trans := helpers.Validate()
	err := valid.Struct(socialMediaRequest)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Error Validation",
			"message": errs.Translate(trans),
		})
		return
	}

	idb.DB.Debug().Table("social_media").Model(&socialMedia).Where("ud = ?", socialMediaId).Updates(models.SocialMedia{
		Name:             socialMediaRequest.Name,
		Social_media_url: socialMediaRequest.Social_media_url,
		Updated_at:       time.Now(),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.Id,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.Social_media_url,
		"user_id":          socialMedia.User_id,
		"updated_at":       socialMedia.Updated_at,
	})
}

func (idb *InDB) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")

	err := idb.DB.Debug().Table("social_media").Where("id = ?", socialMediaId).Delete(models.SocialMedia{}).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Error Delete Social",
			"message": "Cannot find Social Media",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
