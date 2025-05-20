package main

import (
	midl "murweb/internal/middleware"
	"murweb/internal/tools"
	hand "murweb/internal/web/handlers"
	"murweb/repository"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := tools.LoadEnv(); err != nil {
		panic(err)
	}

	file, err := tools.InitLogFile()
	if err != nil {
		panic(err)
	}

	logger, err := tools.InitLogger(file)
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	db, err := repository.NewDBConn(os.Getenv("PG_CONN"), "test", "Users")
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}
	defer db.Close()

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
