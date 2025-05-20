package handlers

import (
	"murweb/internal/models"
	"murweb/internal/service"
	"murweb/internal/tools"
	"murweb/messages"
	"murweb/repository"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CheckTokenHand(g *gin.Context) {
	header := g.GetHeader("Authorization")
	logger := g.MustGet("logger").(*zap.Logger)
	login := g.MustGet("login").(string)
	if !strings.HasPrefix(header, "Bearer ") {
		logger.Error(messages.ErrMissOrInvToken.Error(), zap.String("login", login))
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": messages.ErrMissOrInvToken.Error()})
		return
	}

	token := strings.TrimPrefix(header, "Bearer ")
	claims, err := tools.ValidateJWT(token, tools.KeyFunc)
	if err != nil {
		logger.Error(messages.ErrInvOrExpToken.Error(), zap.String("login", login))
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": messages.ErrInvOrExpToken.Error()})
		return
	}
	g.Set("login", claims["login"])
	g.JSON(http.StatusOK, gin.H{"message": "you are " + claims["login"].(string)})
}

func DepositeHand(db repository.DataBase) gin.HandlerFunc {
	return func(g *gin.Context) {
		logger := g.MustGet("logger").(*zap.Logger)
		login := g.MustGet("login").(string)

		var deposite models.DepositeRequest
		if err := g.ShouldBindJSON(&deposite); err != nil {
			logger.Error(messages.ErrBadConvertReqToJSON.Error(), zap.String("login", login))
			g.JSON(http.StatusBadRequest, models.DepositeResponse{Message: err.Error()})
			return
		}

		err := service.DepositeUser(db, login, deposite.Amount)
		if err != nil {
			logger.Error(err.Error(), zap.String("login", login))
			g.JSON(http.StatusConflict, models.DepositeResponse{Message: err.Error()})
			return
		}
		g.JSON(http.StatusOK, models.DepositeResponse{Message: messages.DeposSucc})
	}

}

func GetBalanceHand(db repository.DataBase) gin.HandlerFunc {
	return func(g *gin.Context) {
		logger := g.MustGet("logger").(*zap.Logger)
		login := g.MustGet("login").(string)
		amount, err := service.GetUserBalans(db, login)
		if err != nil {
			logger.Error(err.Error(), zap.String("login", login))
			g.JSON(http.StatusConflict, models.BalanceResponse{Message: err.Error()})
			return
		}
		logger.Info("getted balance", zap.String("login", login))
		g.JSON(http.StatusOK, models.BalanceResponse{Message: strconv.FormatInt(amount, 10)})
	}
}

func TransferMoneyHand(db repository.DataBase) gin.HandlerFunc {
	return func(g *gin.Context) {
		logger := g.MustGet("logger").(*zap.Logger)
		login := g.MustGet("login").(string)

		var tmpTrans models.TransferRequest
		if err := g.ShouldBindJSON(&tmpTrans); err != nil {
			logger.Error(messages.ErrBadConvertReqToJSON.Error(), zap.String("login", login))
			g.JSON(http.StatusOK, models.TransferResponse{Message: err.Error()})
			return
		}

		if err := service.TransferMoney(db, login, tmpTrans.Login, tmpTrans.Amount); err != nil {
			logger.Error("can't transfer money to "+tmpTrans.Login, zap.String("login", login))
			g.JSON(http.StatusOK, models.TransferResponse{Message: err.Error()})
			return
		}
		logger.Info("Money send succsesfully to "+tmpTrans.Login, zap.String("login", login), zap.String("amount", tmpTrans.Amount))
		g.JSON(http.StatusOK, models.TransferResponse{Message: messages.TransSucc})
	}
}
