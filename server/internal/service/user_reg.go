package service

import (
	"net/http"
	"ubank/internal/models"
	"ubank/internal/tools"
	msg "ubank/messages"
	repos "ubank/repository"
)

func RegUser(db repos.DataBase, user models.RegRequest) (int, error) {
	exists, err := db.CheckExistsUser(user.Email)
	if err != nil {
		return http.StatusConflict, err
	}
	if exists {
		return http.StatusConflict, msg.ErrLoginExists
	}

	pass, err := tools.HashPassword(user.Password)
	if err != nil {
		return http.StatusConflict, msg.ErrHashPass
	}

	if err := db.CreateUser(user.Email, pass); err != nil {
		return http.StatusConflict, err
	}

	return http.StatusOK, nil
}
