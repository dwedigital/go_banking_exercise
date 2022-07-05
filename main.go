package main

import (
	models "banking/models"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	// sort the keys of the map
	var accountNumbers []int
	for k := range models.Customers() {
		accountNumbers = append(accountNumbers, k)
	}
	sort.Ints(accountNumbers)

	// iterate over the sorted keys and print account balances
	for _, k := range accountNumbers {
		fmt.Println(models.Customers()[k].CheckBalance())
	}

	Program()

}

// getInput gets the input from the user
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = formatInput(text)
	return text
}

// formatInput removes the newline character from the input & other formatting
func formatInput(text string) (input string) {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	input = strings.Replace(text, "/n", "", -1)
	return
}

func Program() {
	text := getInput()
	fmt.Println(text)
}
