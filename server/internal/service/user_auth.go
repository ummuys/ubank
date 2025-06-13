package service

import (
	"net/http"
	"ubank/internal/models"
	"ubank/internal/tools"
	msg "ubank/messages"
	repos "ubank/repository"
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
