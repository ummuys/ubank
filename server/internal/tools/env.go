package tools

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return fmt.Errorf("can't load env: %w", err)
	}
	return nil
}
