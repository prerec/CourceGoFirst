package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		v   int
		w   int
		s   string
		arr [2]int
	)

	for i := 0; i < 2; i++ {
		fmt.Scan(&arr[i])
	}
	v = arr[0]
	w = arr[1]
	fmt.Scan(&s)

	input, err := os.Create("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer input.Close()

	for _, num := range arr {
		_, err = fmt.Fprintf(input, "%d\n", num)
		if err != nil {
			fmt.Println("Ошибка при записи числа:", err)
			return
		}
	}
	input.WriteString(s)

	output, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	for i := 0; i < v; i++ {
		for j := 0; j < w; j++ {
			if i == 0 || i == v-1 {
				output.WriteString(s)
				if err != nil {
					fmt.Println("Ошибка при записи числа:", err)
					return
				}
			} else {
				if j == 0 {
					output.WriteString(s)
				}
				if j == w-1 {
					for j := 1; j < w-1; j++ {
						output.WriteString(" ")
					}
					output.WriteString(s)
				}
			}
		}
		output.WriteString("\n")
	}

}
