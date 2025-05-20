package service

import (
	"bufio"
	"fmt"
	"murtest/tools"
	hand "murtest/web/handlers"
	"os"
)

func do() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if _, err := hand.Check(); err != nil {
			fmt.Println("|E| Catched err: ", err.Error())
			return
		}
		tools.DOText()
		scanner.Scan()
		if scanner.Text() == "4" {
			break
		}
		checkAnswerDO(scanner.Text())
	}
}

func checkAnswerDO(ans string) {
	var err error
	var msg string
	switch {
	case ans == "1":
		msg, err = hand.GetBalace()
	case ans == "2":
		depos := tools.ReadDepos()
		msg, err = hand.Deposite(depos)
	case ans == "3":
		transfer := tools.ReadTransfer()
		msg, err = hand.TransferMoney(transfer)
	}
	if err != nil {
		fmt.Printf("\n|E| Catched err: %s\n", err)
	} else if msg != "" {
		fmt.Printf("\n|A| Answer from server: %s\n", msg)
	}
}
