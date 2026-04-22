package main

import (
	"errors"
	"fmt"
)

var val string
var val_chel string
var chislo int

func main() {
	for {
		_, err1 := vvod_ish()
		if err1 != nil {
			fmt.Println("Ошибка: ", err1)
			continue
		}
		break
	}
	for {
		fmt.Println("Введите количество:")
		var err2 error
		_, err2 = fmt.Scan(&chislo)
		if err2 != nil {
			fmt.Println("Ошибка:", err2)
			continue
		}
		break
	}

	for {
		_, err3 := vvod_chel()
		if err3 != nil {
			fmt.Println("Ошибка:", err3)
			continue
		}
		break
	}
	itog := raschet()
	print(itog)

}

func vvod_ish() (string, error) {
	fmt.Println("Введите исходную валюту в формате USD/EUR/RUB:")
	fmt.Scan(&val)
	if val == "USD" || val == "EUR" || val == "RUB" {
		return val, nil
	}
	return val, errors.New("Некорректно введена исходная валюта")
}

func vvod_chel() (string, error) {

	fmt.Println("Введите целевую валюту в формате USD/EUR/RUB:")
	fmt.Scan(&val_chel)
	if val == "USD" && val_chel == "USD" {
		return val_chel, errors.New("Целевая валюта не может быть равна исходной")
	} else if val == "EUR" && val_chel == "EUR" {
		return val_chel, errors.New("Целевая валюта не может быть равна исходной")
	} else if val == "RUB" && val_chel == "RUB" {
		return val_chel, errors.New("Целевая валюта не может быть равна исходной")
	} else if val_chel != "USD" && val_chel != "EUR" && val_chel != "RUB" {
		return val_chel, errors.New("Некорректно введена целевая валюта")
	}
	return val_chel, nil

}

func raschet() float64 {
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
