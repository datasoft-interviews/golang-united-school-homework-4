package string_sum

import (
	"errors"
	"strconv"
	"testing"
)

var revMsg = "StringSum() = %s, want %s, err %s"

func Test1(t *testing.T) {
	input := "3+5"
	want := "8"
	if got, err := StringSum(input); got != want || err != nil {
		t.Errorf(revMsg, got, want, err)
	}
}

func Test2(t *testing.T) {
	input := "-3+5"
	want := "2"
	if got, err := StringSum(input); got != want || err != nil {
		t.Errorf(revMsg, got, want, err.Error())
	}
}

func Test3(t *testing.T) {
	input := "3-5"
	want := "-2"
	if got, err := StringSum(input); got != want || err != nil {
		t.Errorf(revMsg, got, want, err.Error())
	}
}

func Test4(t *testing.T) {
	input := "-3-5"
	want := "-8"
	if got, err := StringSum(input); got != want || err != nil {
		t.Errorf(revMsg, got, want, err.Error())
	}
}

func Test5(t *testing.T) {
	input := "+3-5"
	want := "-2"
	if got, err := StringSum(input); got != want || err != nil {
		t.Errorf(revMsg, got, want, err.Error())
	}
}

func TestWithSpace(t *testing.T) {
	input := "+3 -5"
	want := "-2"
	if got, err := StringSum(input); got != want || err != nil {
		t.Errorf(revMsg, got, want, err.Error())
	}
}

func Test4Empty(t *testing.T) {
	input := ""
	if got, err := StringSum(input); err.Error() != errorEmptyInput.Error() {
		t.Errorf(revMsg, got, "errorEmptyInput", err)
	}
	input = "   "
	if got, err := StringSum(input); err.Error() != errorEmptyInput.Error() {
		t.Errorf(revMsg, got, "errorEmptyInput", err)
	}
}

func Test1Operand(t *testing.T) {
	input := "2"
	if got, err := StringSum(input); err.Error() != errorNotTwoOperands.Error() {
		t.Errorf(revMsg, got, "errorNotTwoOperands", err)
	}
	input = "+2"
	if got, err := StringSum(input); err.Error() != errorNotTwoOperands.Error() {
		t.Errorf(revMsg, got, "errorNotTwoOperands", err)
	}
	input = "-2"
	if got, err := StringSum(input); err.Error() != errorNotTwoOperands.Error() {
		t.Errorf(revMsg, got, "errorNotTwoOperands", err)
	}
	/*
		input = "2+"
		if got, err := StringSum(input); err.Error() != "strconv.Atoi: parsing \"+\": invalid syntax" {
			t.Errorf(revMsg, got, "NumError", err)
		}
		input = "2-"
		if got, err := StringSum(input); err.Error() != "strconv.Atoi: parsing \"-\": invalid syntax" {
			t.Errorf(revMsg, got, "NumError", err)
		}
	*/
}

func Test3Operand(t *testing.T) {
	input := "2+3+4"
	if got, err := StringSum(input); err.Error() != errorNotTwoOperands.Error() {
		t.Errorf(revMsg, got, "errorNotTwoOperands", err)
	}
	input = "+2+3+4"
	if got, err := StringSum(input); err.Error() != errorNotTwoOperands.Error() {
		t.Errorf(revMsg, got, "errorNotTwoOperands", err)
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
