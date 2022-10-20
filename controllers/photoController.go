package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type InDB struct {
	DB *gorm.DB
}

type PhotoRequest struct {
	Title     string `json:"title" gorm:"type varchar(191);not null" validate:"required"`
	Caption   string `json:"caption" gorm:"type varchar(191);not null" validate:"required"`
	Photo_url string `json:"photo_url" gorm:"type varchar(191);not null" validate:"required"`
}

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username" `
	Email    string `json:"email" `
}

type PhotoResponse struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	User_Id    uint      `json:"user_id"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	User       *User     `json:"user"`
}

func (idb *InDB) Create(ctx *gin.Context) {
	var newPhoto PhotoRequest
	err := ctx.ShouldBindJSON(&newPhoto)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	valid, trans := helpers.Validate()
	err = valid.Struct(newPhoto)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Error validator",
			"message": errs.Translate(trans),
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"err":    err,
		})
	}

	var PhotoModel = models.Photo{
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		Photo_url: newPhoto.Photo_url,
		User_id:   userId,
	}

	err = idb.DB.Debug().Create(&PhotoModel).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         PhotoModel.ID,
		"title":      PhotoModel.Title,
		"caption":    PhotoModel.Caption,
		"photo_url":  PhotoModel.Photo_url,
		"user_id":    PhotoModel.User_id,
		"created_at": PhotoModel.Created_at,
	})
}

func (idb *InDB) Get(ctx *gin.Context) {
	var photos []PhotoResponse

	errGetUser := idb.DB.Debug().Table("photos").Preload("User").Find(&photos).Error

	if errGetUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result":  nil,
			"error":   "Error Get Photo",
			"message": errGetUser.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": photos,
	})
}

func (idb *InDB) DeletePhoto(ctx *gin.Context) {

	id := ctx.Param("photoId")
	err := idb.DB.Debug().Table("photos").Where("id = ?", id).Delete(&models.Photo{}).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"photo": nil,
			"err":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}

func (idb *InDB) UpdatePhoto(ctx *gin.Context) {
	var photo PhotoRequest
	var photoModel models.Photo

	id := ctx.Param("id")
	errGetPhoto := idb.DB.Debug().Table("photos").Where("id = ?", id).Take(&photoModel).Error

	if errGetPhoto != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"photo": nil,
			"err":   errGetPhoto,
		})
		return
	}

	err := ctx.ShouldBindJSON(&photo)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"err":    err,
		})
	}

	valid, trans := helpers.Validate()
	err = valid.Struct(photo)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Error validator",
			"message": errs.Translate(trans),
		})
	}

	errUpdate := idb.DB.Debug().Table("photos").Model(&photoModel).Where("id = ?", id).Updates(models.Photo{
		Title:      photo.Title,
		Caption:    photo.Caption,
		Photo_url:  photo.Photo_url,
		Updated_at: time.Now(),
	}).Error

	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    errUpdate.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         photoModel.ID,
		"title":      photoModel.Title,
		"caption":    photoModel.Caption,
		"photo_url":  photoModel.Photo_url,
		"user_id":    photoModel.User_id,
		"updated_at": photoModel.Updated_at,
	})
}
