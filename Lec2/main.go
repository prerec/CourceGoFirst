package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var (
		login string
		email string
	)
	fmt.Scan(&login, &email)
	if (utf8.RuneCountInString(login) < 10) || strings.Contains(login, "@") {
		fmt.Println("Некорректный логин")
		return
	}
	if (strings.Contains(email, "@") != true) || (strings.Contains(email, ".") != true) {
		fmt.Println("Некорректная почта")
		return
	}
	fmt.Println("ОК")
}
