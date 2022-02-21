package main

import (
	"fmt"
	utils "go-homeworks/simple-tasks"
)

func main() {
	x := utils.ScanNum("Начальная сумма вклада: ")
	q := utils.ScanNum("Процентная ставка (%): ")
	r := 1 + q/100
	const n = 5
	for i := 0; i < n; i++ {
		x *= r
	}
	fmt.Printf("Сумма вклада через %d лет: %.2f\n", n, x)
}
