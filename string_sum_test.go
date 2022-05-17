package string_sum

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	input := "3+5"
	want := "8"
	if got, err := StringSum(input); got != want {
		if err != nil {
			fmt.Println(err)
		}
		t.Errorf("ReverseString() = %s, want %s", got, want)
	}
}

func Test2(t *testing.T) {
	input := "-3+5"
	want := "2"
	if got, err := StringSum(input); got != want {
		if err != nil {
			fmt.Println(err)
		}
		t.Errorf("ReverseString() = %s, want %s", got, want)
	}
}

func Test3(t *testing.T) {
	input := "3-5"
	want := "-2"
	if got, err := StringSum(input); got != want {
		if err != nil {
			fmt.Println(err)
		}
		t.Errorf("ReverseString() = %s, want %s", got, want)
	}
}

func Test4(t *testing.T) {
	input := "-3-5"
	want := "-8"
	if got, err := StringSum(input); got != want {
		if err != nil {
			fmt.Println(err)
		}
		t.Errorf("ReverseString() = %s, want %s", got, want)
	}
}

func Test5(t *testing.T) {
	input := "+3-5"
	want := "-2"
	if got, err := StringSum(input); got != want {
		if err != nil {
			fmt.Println(err)
		}
		t.Errorf("ReverseString() = %s, want %s", got, want)
	}
}

func Test4Empty(t *testing.T) {
	input := ""
	if got, err := StringSum(input); err != errorEmptyInput {
		t.Errorf("ReverseString() don't return errorEmptyInput, but return %s", got)
	}
	input = "   "
	if got, err := StringSum(input); err != errorEmptyInput {
		t.Errorf("ReverseString() don't return errorEmptyInput, but return %s", got)
	}
}

func Test1Operand(t *testing.T) {
	input := "2"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf("ReverseString() don't return errorNotTwoOperands, but return %s", got)
	}
	input = "+2"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf("ReverseString() don't return errorNotTwoOperands, but return %s", got)
	}
	input = "-2"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf("ReverseString() don't return errorNotTwoOperands, but return %s", got)
	}
	input = "2+"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf("ReverseString() don't return errorNotTwoOperands, but return %s", got)
	}
	input = "2-"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf("ReverseString() don't return errorNotTwoOperands, but return %s", got)
	}
}

func Test3Operand(t *testing.T) {
	input := "2+3+4"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf("ReverseString() don't return errorNotTwoOperands, but return %s", got)
	}
	input = "+2+3+4"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf("ReverseString() don't return errorNotTwoOperands, but return %s", got)
	}
}

func TestNotValid(t *testing.T) {
	input := "a+2"
	if got, err := StringSum(input); err != firstOperandError {
		t.Errorf("ReverseString() don't return firstOperandError, but return %s, %s", got, err)
	}
	input = "2+a"
	if got, err := StringSum(input); err != secondOperandError {
		t.Errorf("ReverseString() don't return secondOperandError, but return %s, %s", got, err)
	}
}
