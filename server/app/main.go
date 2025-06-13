package main

import (
	"os"
	"ubank/internal/router"
	"ubank/internal/tools"
	"ubank/repository"
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
		logger.Error(err.Error())
		panic(err)
	}
	defer db.Close()

	router.InitRouter(logger, db)
}
