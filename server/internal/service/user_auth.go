package service

import (
	"murweb/internal/models"
	"murweb/internal/tools"
	msg "murweb/messages"
	repos "murweb/repository"
	"net/http"
)

func AuthUser(db repos.DataBase, user models.AuthRequest) (int, string, error) {
	exists, err := db.CheckExistsUser(user.Email)
	if err != nil {
		return http.StatusConflict, "", err
	}
	if !exists {
		return http.StatusConflict, "", msg.ErrLoginNotExists
	}

	if pass, err := db.GetPassword(user.Email); err == nil {
		if tools.CheckHash(user.Password, pass) {
			token, err := tools.GenerateJWT(user.Email, "ummuys")
			if err != nil {
				return http.StatusBadRequest, "", msg.ErrCreateJWT
			}
			return http.StatusOK, token, nil
		} else {
			return http.StatusForbidden, "", msg.ErrIncorrectPass
		}
	} else {
		return http.StatusForbidden, "", err
	}
}
