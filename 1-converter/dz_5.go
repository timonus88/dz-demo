package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var val string
var val_chel string
var chislo int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		_, err1 := vvod_ish(scanner)
		if err1 != nil {
			fmt.Println("Ошибка: ", err1)
			continue
		}
		break
	}

	for {
		fmt.Println("Введите количество:")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			fmt.Println("Пустой ввод. Попробуйте снова.")
			continue
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите целое число.")
			continue
		}
		if num < 0 {
			fmt.Println("Ошибка: число должно быть неотрицательным.")
			continue
		}
		chislo = num
		break
	}

	for {
		_, err3 := vvod_chel(scanner)
		if err3 != nil {
			fmt.Println("Ошибка:", err3)
			continue
		}
		break
	}

	itog := raschet(chislo, val, val_chel)
	fmt.Println(itog)
}

func vvod_ish(scanner *bufio.Scanner) (string, error) {
	fmt.Println("Введите исходную валюту в формате USD/EUR/RUB:")
	scanner.Scan()
	val = strings.TrimSpace(scanner.Text())
	if val == "USD" || val == "EUR" || val == "RUB" {
		return val, nil
	}
	return val, errors.New("Некорректно введена исходная валюта")
}

func vvod_chel(scanner *bufio.Scanner) (string, error) {
	fmt.Println("Введите целевую валюту в формате USD/EUR/RUB:")
	scanner.Scan()
	val_chel = strings.TrimSpace(scanner.Text())
	if val == val_chel {
		return val_chel, errors.New("Целевая валюта не может быть равна исходной")
	}
	if val_chel != "USD" && val_chel != "EUR" && val_chel != "RUB" {
		return val_chel, errors.New("Некорректно введена целевая валюта")
	}
	return val_chel, nil
}

func raschet(chislo int, val string, val_chel string) float64 {
	var a float64
	const usd_eur = 0.8564
	const usd_rub = 76.97

	switch {
	case val == "USD" && val_chel == "EUR":
		a = usd_eur * float64(chislo)
	case val == "USD" && val_chel == "RUB":
		a = usd_rub * float64(chislo)
	case val == "EUR" && val_chel == "RUB":
		a = (usd_rub / usd_eur) * float64(chislo)
	case val == "EUR" && val_chel == "USD":
		a = (1 / usd_eur) * float64(chislo)
	case val == "RUB" && val_chel == "USD":
		a = (1 / usd_rub) * float64(chislo)
	case val == "RUB" && val_chel == "EUR":
		a = ((1 / usd_rub) * usd_eur) * float64(chislo)
	}
	return a
}
