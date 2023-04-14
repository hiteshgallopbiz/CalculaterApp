package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	exp := 8
	res := Addition(5, 3)
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

func BenchmarkAddition(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Addition(1, 2)
	}
}

func BenchmarkSubstraction(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Substraction(1, 2)
	}
}

func BenchmarkMultiplication(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Multiplication(1, 2)
	}
}

func BenchmarkDivision(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Division(9, 3)
	}
}
