package string_sum

import (
	"errors"
	"strconv"
	"testing"
)

var (
	revMsg  = "ReverseString() = %s, want %s"
	revMsg2 = "ReverseString() don't return errorEmptyInput, but return %s"
	revMsg3 = "ReverseString() don't return errorNotTwoOperands, but return %s"
)

func Test1(t *testing.T) {
	input := "3+5"
	want := "8"
	if got, _ := StringSum(input); got != want {
		t.Errorf(revMsg, got, want)
	}
}

func Test2(t *testing.T) {
	input := "-3+5"
	want := "2"
	if got, _ := StringSum(input); got != want {
		t.Errorf(revMsg, got, want)
	}
}

func Test3(t *testing.T) {
	input := "3-5"
	want := "-2"
	if got, _ := StringSum(input); got != want {
		t.Errorf(revMsg, got, want)
	}
}

func Test4(t *testing.T) {
	input := "-3-5"
	want := "-8"
	if got, _ := StringSum(input); got != want {
		t.Errorf(revMsg, got, want)
	}
}

func Test5(t *testing.T) {
	input := "+3-5"
	want := "-2"
	if got, _ := StringSum(input); got != want {
		t.Errorf(revMsg, got, want)
	}
}

func Test4Empty(t *testing.T) {
	input := ""
	if got, err := StringSum(input); err != errorEmptyInput {
		t.Errorf(revMsg2, got)
	}
	input = "   "
	if got, err := StringSum(input); err != errorEmptyInput {
		t.Errorf(revMsg2, got)
	}
}

func Test1Operand(t *testing.T) {
	input := "2"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf(revMsg3, got)
	}
	input = "+2"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf(revMsg3, got)
	}
	input = "-2"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf(revMsg3, got)
	}
	input = "2+"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf(revMsg3, got)
	}
	input = "2-"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf(revMsg3, got)
	}
}

func Test3Operand(t *testing.T) {
	input := "2+3+4"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf(revMsg3, got)
	}
	input = "+2+3+4"
	if got, err := StringSum(input); err != errorNotTwoOperands {
		t.Errorf(revMsg3, got)
	}
}

func TestNotValid(t *testing.T) {
	input := "a+2"
	want := &strconv.NumError{}
	if got, err := StringSum(input); errors.Is(err, want) {
		t.Errorf("ReverseString() don't return firstOperandError, but return %s, %v", got, err)
	}
	input = "2+a"
	if got, err := StringSum(input); errors.Is(err, want) {
		t.Errorf("ReverseString() don't return secondOperandError, but return %s, %v", got, err)
	}
}
