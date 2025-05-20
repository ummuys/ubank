package service

import (
	"errors"
	"fmt"
	"murweb/repository"
	"strconv"
)

func DepositeUser(db repository.DataBase, login string, amount string) error {
	amountInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return fmt.Errorf("can't convert cash to int64: %w", err)
	}

	if err := db.Deposite(login, amountInt); err != nil {
		return err
	}
	return nil
}

func GetUserBalans(db repository.DataBase, login string) (int64, error) {
	cash, err := db.GetUserBalans(login)
	if err != nil {
		return 0, err
	}
	return cash, nil
}

func TransferMoney(db repository.DataBase, loginFrom, loginTo string, amount string) error {
	amountInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return fmt.Errorf("can't convert amounr to int64: %w", err)
	}

	currCash, err := db.GetUserBalans(loginFrom)
	if err != nil {
		return err
	}

	if currCash-amountInt < 0 {
		return errors.New("Unable to complete transaction: balance is insufficient")
	}

	if exists, err := db.CheckExistsUser(loginTo); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("%s doen't exists", loginTo)
	}

	if err := db.TransferMoney(loginFrom, loginTo, amountInt); err != nil {
		return err
	}
	return nil
}
