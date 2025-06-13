package router

import (
	midl "ubank/internal/middleware"
	hand "ubank/internal/web/handlers"
	"ubank/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(logger *zap.Logger, db repository.DataBase) {
	g := gin.Default()
	g.Use(midl.AddZapLogger(logger))

	login := g.Group("/me")
	login.Use(midl.JWTRequest())
	login.GET("", hand.CheckTokenHand)
	login.POST("/deposite", hand.DepositeHand(db))
	login.GET("/balance", hand.GetBalanceHand(db))
	login.POST("/transfer", hand.TransferMoneyHand(db))
	g.POST("/auth", hand.Auth(db))
	g.POST("/reg", hand.Reg(db))
	g.Run()
}
