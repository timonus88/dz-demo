package main

import "fmt"

func main() {
	const usd_eur = 0.8564
	const usd_rub = 76.97
	var eur_rub = usd_rub / usd_eur
	fmt.Print("EUR в RUB = ", eur_rub)
}
