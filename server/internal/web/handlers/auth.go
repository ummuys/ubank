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

func Auth(db repository.DataBase) gin.HandlerFunc {
	return func(g *gin.Context) {
		logger := g.MustGet("logger").(*zap.Logger)

		var user models.AuthRequest
		if err := g.ShouldBindJSON(&user); err != nil {
			logger.Error("can't convert to json")
			g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		status, token, err := service.AuthUser(db, user)
		if err != nil {
			logger.Error(err.Error(), zap.String("login", user.Email))
			g.JSON(status, models.AuthResponse{Message: err.Error()})
			return
		}

		g.JSON(status, models.AuthResponse{Message: messages.AuthSucc, Token: token})
		logger.Info("token was created", zap.String("login", user.Email))
	}
}
