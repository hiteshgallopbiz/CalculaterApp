package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	exp := 8
	res := .Addition()
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}
func TestSubstraction(t *testing.T) {
	exp := 8
	res := Substraction(11, 3)
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}

func TestMultiplication(t *testing.T) {
	exp := 33
	res := Multiplication(11, 3)
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}

func TestDivision(t *testing.T) {
	exp := 3
	res := Division(9, 3)
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}
