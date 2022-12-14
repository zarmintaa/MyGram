package controllers

import (
	"final-project/dto"
	"final-project/helpers"
	"final-project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/url"
	"time"
)

func (idb *InDB) CreateSocialMedia(ctx *gin.Context) {
	var SocialMedia models.SocialMedia
	var SocialMediaReq dto.SocialMediaRequest

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

	_, err = url.ParseRequestURI(SocialMediaReq.SocialMediaUrl)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid type url = ",
		})
		return
	}

	SocialMedia = models.SocialMedia{
		Name:           SocialMediaReq.Name,
		SocialMediaUrl: SocialMediaReq.SocialMediaUrl,
		CreatedAt:      time.Now(),
		UserId:         uint(userData["id"].(float64)),
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
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserId,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func (idb *InDB) GetSocialMedia(ctx *gin.Context) {
	var listSocialMedia []dto.SocialMediaResponse
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	err := idb.DB.Debug().Table("social_media").Where("user_id = ?", userId).Preload("User").Find(&listSocialMedia).Error

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
	var socialMediaRequest dto.SocialMediaRequest
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

	errGet := idb.DB.Debug().Table("social_media").Model(&socialMedia).Where("id = ?", socialMediaId).Take(&socialMedia).Error

	if errGet != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Error Update",
			"message": "Unable to find Photo with id = " + socialMediaId,
		})
		return
	}

	errJson = idb.DB.Debug().Table("social_media").Model(&socialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{
		Name:           socialMediaRequest.Name,
		SocialMediaUrl: socialMediaRequest.SocialMediaUrl,
		UpdatedAt:      time.Now(),
	}).Error

	if errJson != nil {
		ctx.JSON(http.StatusBadRequest, errJson.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.Id,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserId,
		"updated_at":       socialMedia.UpdatedAt,
	})
}

func (idb *InDB) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")

	errGet := idb.DB.Debug().Table("social_media").Where("id = ?", socialMediaId).Take(models.SocialMedia{}).Error

	if errGet != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot find Social media with id = " + socialMediaId,
			"error":   errGet.Error(),
		})
		return
	}

	err := idb.DB.Debug().Table("social_media").Where("id = ?", socialMediaId).Delete(models.SocialMedia{}).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Error Delete Social",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
