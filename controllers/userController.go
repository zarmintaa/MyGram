package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username" gorm:"type varchar(10);unique;not null" validate:"required,lte=100"`
	Email    string `json:"email" gorm:"type varchar(191);not null;unique" validate:"required,lte=100"`
	Password string `json:"password"  gorm:"type varchar(191); not null" validate:"required,lte=100"`
	Age      int    `json:"age" validate:"required,lte=100,gte=8"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type varchar(191);not null;unique" validate:"required,lte=100,gte=8"`
	Password string `json:"password"  gorm:"type varchar(191); not null" validate:"required,lte=100,gte=8"`
}

func (idb *InDB) Register(ctx *gin.Context) {
	var newUser RegisterRequest
	err := ctx.ShouldBindJSON(&newUser)

	valid, trans := helpers.Validate()
	err = valid.Struct(newUser)

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

	checkEmailFormat := helpers.EmailFormatValidation(newUser.Email)

	if !checkEmailFormat {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email field must a valid format!",
		})
		return
	}

	var User = models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: helpers.HashPassword(newUser.Password),
		Age:      newUser.Age,
	}

	err = idb.DB.Debug().Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.Id,
		"username": User.Username,
	})

}

func (idb *InDB) Login(ctx *gin.Context) {
	var userReq LoginRequest
	err := ctx.ShouldBindJSON(&userReq)
	valid, trans := helpers.Validate()
	err = valid.Struct(userReq)

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

	checkEmailFormat := helpers.EmailFormatValidation(userReq.Email)

	if !checkEmailFormat {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email field must a valid format!",
		})
		return
	}

	var UserModel = models.User{
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	err = idb.DB.Debug().Where("email = ?", UserModel.Email).Take(&UserModel).Error

	errHash := helpers.ComparePass(userReq.Password, UserModel.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email not found",
		})
		return
	}

	if !errHash {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong Password!",
			"error":   errHash,
		})
		return
	}

	token := helpers.GenerateToken(UserModel.Id, UserModel.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (idb *InDB) UpdateUser(ctx *gin.Context) {
	var userRequest LoginRequest
	err := ctx.ShouldBindJSON(&userRequest)

	valid, trans := helpers.Validate()
	err = valid.Struct(userRequest)

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

	var UserModel = models.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	err = idb.DB.Debug().Where("email = ?", UserModel.Email).First(&UserModel).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email not found",
		})
		return
	}

	err = idb.DB.Debug().Model(&models.User{}).Where("email = ?", userRequest.Email).Update("password", userRequest.Password).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": "Password failed to update",
			"error":  err,
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        UserModel.Id,
		"email":     UserModel.Email,
		"username":  UserModel.Username,
		"age":       UserModel.Age,
		"update_at": UserModel.Updated_at,
	})
}

func (idb *InDB) DeleteUser(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	var userModel models.User

	errGetUser := idb.DB.Debug().First(&userModel, "id = ?", userId).Error

	if errGetUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}
	errDelete := idb.DB.Debug().Where("email = ?", userModel.Email).Delete(&userModel).Error

	if errDelete != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong!",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
