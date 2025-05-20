package tools

import (
	"bufio"
	"fmt"
	"murtest/models"
	"os"
)

func ReadInfo() *models.User {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nВведите свой логин:")
	scanner.Scan()
	login := scanner.Text()

	fmt.Println("Введите свой пароль:")
	scanner.Scan()
	pass := scanner.Text()
	return &models.User{Login: login, Pass: pass}
}

func ReadDepos() *models.DepositeRequest {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nСколько желаете положить на баланс:")
	scanner.Scan()
	amount := scanner.Text()
	return &models.DepositeRequest{Amount: amount}
}

func ReadTransfer() *models.TransferRequest {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nКому вы хотите отправить деньги:")
	scanner.Scan()
	loginTo := scanner.Text()

	fmt.Println("\nКакую сумму вы хотите отправить:")
	scanner.Scan()
	amount := scanner.Text()
	return &models.TransferRequest{Login: loginTo, Amount: amount}
}

func MenuText() {
	fmt.Println("\n||\\|| Что вы хотите сделать? ||\\||\n\n" +
		"1) Зарегистрироваться/Войти\n" +
		"2) Меню банка \n" +
		"3) Выйти \n",
	)
}

func LAText() {
	fmt.Println("\n||\\|| Выберите одно из перечисленных вариантов:||\\||\n\n" +
		"1) Зарегистрироваться\n" +
		"2) Войти в аккаунт\n" +
		"3) Проверить кто ты\n" +
		"4) Вернуться обратно\n",
	)
}

func DOText() {
	fmt.Println("\n||\\|| Выберите одно из перечисленных вариантов:||\\||\n\n" +
		"1) Узнать баланс\n" +
		"2) Внести депозит\n" +
		"3) Отправить деньги\n" +
		"4) Вернуться обратно\n",
	)
}
