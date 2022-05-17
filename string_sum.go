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
	firstOperandError   = fmt.Errorf("First operand is not valid")
	secondOperandError  = fmt.Errorf("Second operand is not valid")
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
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", errorEmptyInput
	}
	del := -1
	for i := 1; i < len(input); i++ {
		if string(input[i]) == "-" || string(input[i]) == "+" {
			if del != -1 {
				return "", errorNotTwoOperands
			}
			del = i
		}
	}
	if del == -1 {
		return "", errorNotTwoOperands
	}
	first_oper := input[0:del]
	second_oper := input[del:]
	if len(second_oper) == 1 {
		return "", errorNotTwoOperands
	}
	first_sum, err := strconv.Atoi(first_oper)
	if err != nil {
		return "", firstOperandError
	}
	second_sum, err := strconv.Atoi(second_oper)
	if err != nil {
		return "", secondOperandError
	}
	return strconv.Itoa(first_sum + second_sum), nil
}
