package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"murtest/models"
	"murtest/repository"
	"net/http"
)

func Check() (string, error) {

	req, err := http.NewRequest("GET", "http://localhost:8080/me", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+repository.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
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

func Deposite(depos *models.DepositeRequest) (string, error) {

	jsonDepos, _ := json.Marshal(depos)
	req, err := http.NewRequest("POST", "http://localhost:8080/me/deposite", bytes.NewBuffer(jsonDepos))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+repository.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
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

func GetBalace() (string, error) {

	req, err := http.NewRequest("GET", "http://localhost:8080/me/balance", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+repository.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if resp.StatusCode == http.StatusOK {
		return result["message"].(string), nil
	}

	return "", errors.New(result["message"].(string))
}

func TransferMoney(transfer *models.TransferRequest) (string, error) {
	jsonTrans, _ := json.Marshal(transfer)
	req, err := http.NewRequest("POST", "http://localhost:8080/me/transfer", bytes.NewBuffer(jsonTrans))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+repository.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if resp.StatusCode == http.StatusOK {
		return result["message"].(string), nil
	}

	return "", errors.New(result["message"].(string))
}
