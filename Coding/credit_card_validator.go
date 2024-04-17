package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validateCreditCard(cardNumber string) bool {

	cardNumber = strings.ReplaceAll(cardNumber, "-", "")
	if _, err := strconv.Atoi(cardNumber); err != nil {
		return false
	}

	if !(strings.HasPrefix(cardNumber, "4") ||
		strings.HasPrefix(cardNumber, "5") ||
		strings.HasPrefix(cardNumber, "6")) {
		return false
	}

	if len(cardNumber) != 16 {
		return false
	}

	var cardRegex = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?)$`)

	return cardRegex.MatchString(cardNumber)

}

func main() {
	var cardCount int
	fmt.Print("Enter the number of cards to be checked ->")
	fmt.Scan(&cardCount)

	for i := 0; i < cardCount; i++ {
		var cardNumber string
		var cardCount = i + 1
		fmt.Printf("Enter card number %d ->", cardCount)
		fmt.Scan(&cardNumber)

		if validateCreditCard(cardNumber) {
			fmt.Printf("This credit card with the number - %s - is valid\n", cardNumber)
		} else {
			fmt.Printf("This credit card with the number - %s - is Invalid\n", cardNumber)
		}

	}
}
