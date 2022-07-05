package main

import (
	models "banking/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// global variables
var customers = models.Customers()

func main() {
	// TODO create a separate function to sort accounts and return the sorted map
	var accountNumbers []int
	// sort the keys of the map
	for k := range customers {
		accountNumbers = append(accountNumbers, k)
	}
	sort.Ints(accountNumbers)

	// iterate over the sorted keys and print account balances
	for _, k := range accountNumbers {
		fmt.Println(customers[k].CheckBalance())
	}

	customers[123].Withdraw(100.00)
	fmt.Println(customers[123].CheckBalance())

	Program()

}

// getInput gets the input from the user
func getInput(helpText string) string {
	fmt.Printf("%s ", helpText)
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
	text := getInput("Enter your account number:")
	account, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(models.Customers()[account].CheckBalance())
}
