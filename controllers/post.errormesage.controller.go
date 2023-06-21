package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mirr254/debughq/models"
	"gorm.io/gorm"
)

type ErrorMessageController struct {
	DB *gorm.DB
}

func NewPostErrorMessageController(DB *gorm.DB) ErrorMessageController {
	return ErrorMessageController{DB}
}

// [...] Create ErrorMessage Handler
func (pemc *ErrorMessageController) CreateErrorMessage( ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.Users)

	var payload *models.PostErrorMessageRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newErrorMessage := models.ErrorMessage {
       ErrorMessage: payload.ErrorMessage,
	   CreatedAt: now,
	   UpdatedAt: now,
	   User: currentUser.ID,
	}

	result := pemc.DB.Create(&newErrorMessage)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key"){
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "An error log with that title already exist"} )
			return 
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()} )
		return 
	} 
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newErrorMessage})
}

//[...] Update ErrorMessage Handler
func (pem *ErrorMessageController) UpdateErrorMessage( ctx *gin.Context) {
	errorMessageId := ctx.Param("errorMessageId")
	currentUser := ctx.MustGet("currentUser").(models.Users)

	var payload *models.UpdateErrorMessage

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error() })
		return 
	}
	
	var updatedErrorMessage models.ErrorMessage
	result := pem.DB.First( &updatedErrorMessage, "id = ?", errorMessageId )

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Error log with that Id exists"})
		return 
	}

	now := time.Now()
	errorMessageToUpdadte := models.ErrorMessage{
	   ErrorMessage: payload.ErrorMessage,
	   CreatedAt:    updatedErrorMessage.CreatedAt,
	   UpdatedAt:    now,
	   User:         currentUser.ID,
	}

	pem.DB.Model(&updatedErrorMessage).Updates(errorMessageToUpdadte)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": updatedErrorMessage})
}

//[....]Get single ErrorMessage
func (pem *ErrorMessageController) GetErrorMessage(ctx *gin.Context) {
	errorMessageId := ctx.Param("errorMessageId")

	var errorMessage models.ErrorMessage
	result := pem.DB.First(&errorMessage, "id = ?", errorMessageId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status":"fail", "message": errorMessage})
		return
	}

}

//[...] Get all the ErrorMessage
func (pem *ErrorMessageController) GetAllErrorMessage(ctx *gin.Context){

	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offSet := (intPage - 1) * intLimit

	var errorMessages []models.ErrorMessage
	results := pem.DB.Limit(intLimit).Offset(offSet).Find(&errorMessages)

	if results.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"error", "meesage": results.Error})
	}

	ctx.JSON(http.StatusOK, gin.H{"status":"success", "results": len(errorMessages), "data": errorMessages})

}

//[...] Delete an ErrorMessage
func (pem *ErrorMessageController) DeleteAnErrorMessage(ctx *gin.Context) {
	errorMessageId := ctx.Param("errorMessageId")
    
	results := pem.DB.Delete(&models.ErrorMessage{}, "id = ?", errorMessageId)
	if results.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message":"ErrorMessage no found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)


}