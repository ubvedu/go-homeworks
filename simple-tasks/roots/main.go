package main

import (
	"fmt"
	utils "go-homeworks/simple-tasks"
	"math"
)

func main() {
	fmt.Println("ax2 + bx + c = 0")
	a := utils.ScanNum("a = ")
	b := utils.ScanNum("b = ")
	c := utils.ScanNum("c = ")
	if a == 0 {
		fmt.Printf("x = %.4g\n", -c/b)
		return
	}
	d := b*b - 4*a*c
	if d == 0 {
		fmt.Printf("x = %.4g\n", -b/2/a)
		return
	}
	if d < 0 {
		fmt.Println("no roots")
		return
	}
	fmt.Printf("x1 = %.4g, x2 = %.4g\n", (-b-math.Sqrt(d))/2/a, (-b+math.Sqrt(d))/2/a)
}
