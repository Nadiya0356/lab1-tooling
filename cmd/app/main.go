package main

import (
	"fmt"

	"github.com/Nadiya0356/lab1-tooling/internal"
)

func main() {
	fmt.Println("Add:", internal.Add(5, 3))

	result, err := internal.Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide:", result)
	}

	// Виправлено: помилка тепер "оброблена"
	_, _ = internal.Divide(5, 0)
}
