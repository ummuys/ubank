package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"murtest/models"
	"net/http"
)

func Reg(user *models.User) (string, error) {
	jsonUser, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8080/reg", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if resp.StatusCode == http.StatusOK {
		return result["message"].(string), nil
	}

	return "", errors.New(result["message"].(string))
}
