package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mirr254/debughq/controllers"
)

type ErrorMessageRouteController struct {
	errorMessageController controllers.ErrorMessageController
}

func NewRouteErrorMessageController ( errorMessageController controllers.ErrorMessageController ) ErrorMessageRouteController {
	return ErrorMessageRouteController{errorMessageController}
}

func (pem *ErrorMessageRouteController) ErrorMessageRoute(rg *gin.RouterGroup) {
	router := rg.Group("errorMessages")

	router.POST("/", pem.errorMessageController.CreateErrorMessage)
	router.GET("/", pem.errorMessageController.GetAllErrorMessage)
	router.PUT("/:errorMessageId", pem.errorMessageController.UpdateErrorMessage)
	router.DELETE("/:errorMessageId", pem.errorMessageController.DeleteAnErrorMessage)
	router.GET("/:errorMessageId",pem.errorMessageController.GetErrorMessage)
}