package service

import (
	"bufio"
	"fmt"
	"murtest/tools"
	hand "murtest/web/handlers"
	"os"
)

func la() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		tools.LAText()
		scanner.Scan()
		if scanner.Text() == "4" {
			break
		}
		checkAnswerLa(scanner.Text())
	}
}

func checkAnswerLa(ans string) {
	var err error
	var msg string
	switch {
	case ans == "1":
		user := tools.ReadInfo()
		msg, err = hand.Reg(user)
	case ans == "2":
		user := tools.ReadInfo()
		msg, err = hand.Auth(user)
	case ans == "3":
		msg, err = hand.Check()
	}
	if err != nil {
		fmt.Printf("\n|E| Catched err: %s\n", err)
	} else {
		fmt.Printf("\n|A| Answer from server: %s\n", msg)
	}
}
