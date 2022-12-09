// Package isbn provides functions to verify ISBN
package main

import "strings"

var replacer = strings.NewReplacer("-", "")

// IsValidISBN checks whether the input is valid ISBN 10 or 13.
func IsValidISBN(input string) bool {
	normalizedInput := replacer.Replace(input)
	if len(normalizedInput) == 13 {
		return isValidISBN13(normalizedInput)
	} else if len(normalizedInput) == 10 {
		return isValidISBN10(normalizedInput)
	}
	return false
}

// IsValidISBN10 check ISBN10 valid.
func isValidISBN10(normalizedInput string) bool {
	var res, charValue, num int

	for index := 0; index < 9; index++ {
		num = int(normalizedInput[index])

		if !(num > 47 && num < 58) {
			return false
		}
		charValue = num - 48
		res += (10 - index) * charValue
	}
	num = int(normalizedInput[9])
	if !((num > 47 && num < 58) || num == 88) {
		return false
	}
	if num == 88 {
		charValue = 10
	} else {
		charValue = num - 48
	}
	res += charValue
	return res%11 == 0
}

// IsValidISBN13 check ISBN13 valid.
func isValidISBN13(normalizedInput string) bool {
	var res, charValue int
	var multiplier int

	for index := 0; index < 12; index++ {
		num := int(normalizedInput[index])
		charValue = num - 48
		if !(num > 47 && num < 58) {
			return false
		}

		if index%2 == 1 {
			multiplier = 3
		} else {
			multiplier = 1
		}
		res += multiplier * charValue
	}

	// check last digit
	charValue = int(normalizedInput[12]) - 48
	if charValue > 9 {
		return false
	}
	remainder := 10 - res%10
	isValid := (remainder < 10 && remainder == charValue)
	isValid = isValid || (remainder == 10 && charValue == 0)
	return isValid
}
