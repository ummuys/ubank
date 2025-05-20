package handlers

import (
	"murweb/internal/models"
	"murweb/internal/service"
	"murweb/messages"
	"murweb/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Reg(db repository.DataBase) gin.HandlerFunc {
	return func(g *gin.Context) {
		logger := g.MustGet("logger").(*zap.Logger)
		var user models.RegRequest
		if err := g.ShouldBindJSON(&user); err != nil {
			logger.Error("can't convert to json")
			g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		status, err := service.RegUser(db, user)
		if err != nil {
			logger.Error(err.Error(), zap.String("login", user.Email), zap.Int("code", status))
			g.JSON(status, models.RegResponse{Message: messages.ErrLoginExists.Error()})
			return
		}
		logger.Info(messages.RegSucc, zap.String("login", user.Email), zap.Int("code", status))
		g.JSON(status, models.RegResponse{Message: messages.RegSucc})
	}
}
