package service

import (
	"murweb/internal/models"
	"murweb/internal/tools"
	msg "murweb/messages"
	repos "murweb/repository"
	"net/http"
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
