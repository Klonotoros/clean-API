package controller

import (
	"clean-API/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoController interface {
	Info(*gin.Context)
}

type infoController struct{}

func newInfoController() InfoController {
	return &infoController{}
}

func (infoController) Info(context *gin.Context) {
	context.JSON(http.StatusOK, dto.ServerInfo{Healthy: true})
}
