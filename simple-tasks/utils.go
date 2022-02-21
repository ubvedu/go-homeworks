package simple_tasks

import "fmt"

func ScanNum(prompt string) float64 {
	fmt.Print(prompt)
	var n float64
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}
	return n
}
