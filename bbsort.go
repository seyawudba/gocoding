package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 10 numbers separated by spaces:\n")

	re := scanner.Scan()

	status, sli := handleUserInput(re, scanner)

	if status {
		BubbleSort(sli)
		fmt.Printf("Your sorted numbers: %v\n", sli)
	}
}

func BubbleSort(sli []int) {
	n := len(sli)
	var i int
	for i = 0; i < n-1; i++ {
		Swap(i, sli)
	}
}

func Swap(i int, sli []int) {
	n := len(sli)
	var j int
	for j = 0; j < n-i-1; j++ {
		if sli[j] > sli[j+1] {
			sli[j], sli[j+1] = sli[j+1], sli[j]
		}

	}
}

func handleUserInput(re bool, scanner *bufio.Scanner) (bool, []int) {
	const capacity = 10
	sli := make([]int, 0, capacity)
	status := false

	if re {
		line := scanner.Text()
		string_parts := strings.Fields(line)

		if len(string_parts) != 10 {
			fmt.Printf("Length of entry must be 10")
			return status, sli
		}
		status, sli = convertInput(string_parts, sli, status)
		return status, sli

	} else {
		fmt.Println("An unexpected error occured!\n")
		return status, sli
	}
}

func convertInput(string_parts []string, sli []int, status bool) (bool, []int) {
	for _, s := range string_parts {
		num, err := strconv.Atoi(s)
		if err == nil {
			sli = append(sli, num)
		} else {
			fmt.Println("An error occured during conversion.\n")
			return status, sli
		}

	}
	status = true
	return status, sli
}
