package service

import (
	"bufio"
	"murtest/tools"
	"os"
)

func Menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		tools.MenuText()
		scanner.Scan()
		if scanner.Text() == "3" {
			return
		}
		checkAnswerMenu(scanner.Text())
	}
}

func checkAnswerMenu(ans string) {
	switch {
	case ans == "1":
		la()
	case ans == "2":
		do()
	}
}
