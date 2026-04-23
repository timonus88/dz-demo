package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		operation, err := inputOperation()
		if err != nil {
			fmt.Println("Ошибка: ", err)
			continue
		}

		slice := inputSlice()
		result := calculateOperation(operation, slice)
		fmt.Printf("Результат (%s): %.2f\n", operation, result)
		break
	}
}

func inputOperation() (string, error) {
	var operation string
	fmt.Println("Введите название операции AVG-среднее или SUM-сумму или MED-медиану:")
	fmt.Scan(&operation)
	if operation == "AVG" || operation == "SUM" || operation == "MED" {
		return operation, nil
	}
	return operation, errors.New("Некорректно введена операция. Повторите ввод.")
}

func inputSlice() []int {
	var number string
	var tr1 []int
	fmt.Printf("Введите целые числа через запятую:\n")
	fmt.Scan(&number)
	parts := strings.Split(number, ",")
	for _, part := range parts {
		trimmed := strings.TrimSpace(part) // убираем пробелы вокруг
		if trimmed == "" {
			continue // пропускаем пустые
		}
		number, err := strconv.Atoi(trimmed)
		if err != nil {
			fmt.Printf("'%s' — Это не целое число, пропускаем шаг\n", trimmed)
			continue // или можно завершить с ошибкой
		}
		tr1 = append(tr1, number) // добавляем в слайс
	}
	return tr1

}

func calculateOperation(operation string, tr1 []int) float64 {
	itogo := 0.0
	if len(tr1) == 0 {
		return 0.0
	}
	if operation == "SUM" {
		for _, number := range tr1 {
			itogo += float64(number)
		}
		return float64(itogo)
	}
	if operation == "AVG" {
		for _, number := range tr1 {
			itogo += float64(number) / float64(len(tr1))
		}
		return float64(itogo)
	}
	if operation == "MED" {
		sorted := make([]int, len(tr1))
		copy(sorted, tr1)
		sort.Ints(sorted)
		n := len(sorted)
		if n%2 == 1 {
			return float64(sorted[n/2])
		} else {
			return float64(sorted[n/2-1]+sorted[n/2]) / 2.0
		}

	}
	return itogo
}
