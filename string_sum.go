package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	ops := make([]string, 0)
	input = strings.TrimSpace(input) // remove spaces at the begin and at the end
	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	buf := string(input[0])
	for i := 1; i < len(input); i++ {
		// remove spaces in the middle
		if string(input[i]) == " " {
			continue
		}
		// new operand starts with sign
		if string(input[i]) == "-" || string(input[i]) == "+" {
			ops = append(ops, buf) // add buffered operand
			buf = string(input[i]) // begin to collect new operand
		} else { // collect operand
			buf = buf + string(input[i])
		}
	}
	// add last operand if it exist
	if len(buf) > 0 {
		ops = append(ops, buf)
	}
	// no operands
	if len(ops) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	// wrong number of operands
	if len(ops) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	firstSum, err := strconv.Atoi(ops[0])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	secondSum, err := strconv.Atoi(ops[1])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return strconv.Itoa(firstSum + secondSum), nil
}
