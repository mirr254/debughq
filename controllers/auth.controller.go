package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mirr254/debughq/models"
	"github.com/mirr254/debughq/utils"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController{
	return AuthController{DB}
}

//[...] Signup User
func (ac *AuthController) SignupUser(ctx *gin.Context) {
	var payload *models.SignupInput

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"failed", "message": err.Error()})
		return 
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"failed", "message": "passwords don't match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"fail", "message": err.Error()})
		return 
	}

	now := time.Now()

	newUser := models.User{
		Name: payload.Name,
		Email: strings.ToLower(payload.Email),
		Password: hashedPassword,
		Role: "user",
		Provider: "default",
		Verified: true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	results := ac.DB.Create(&newUser)

	if results.Error != nil && strings.Contains(results.Error.Error(), "duplicate key value violates unique") {
       ctx.JSON(http.StatusConflict, gin.H{"status":"failed", "message": "A user with that email or name exists"})
	   return 
	} else if results.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"failed", "message": results.Error.Error()})
	}

	userResponse := models.UserResponse{
		Name: newUser.Name,
		ID: newUser.ID,
		Email: newUser.Email,
		Role: newUser.Role,
		Provider: newUser.Provider,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}