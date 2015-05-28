package main

import (
	"testing"
)

func Test_4thPower(t *testing.T) {
	num := writeableAsSumOfNthPower(4)
	if sum(num) != 19316 {
		t.Errorf("sum = %v, want %v", sum(num), 19316)
	}
}

func Test_5thPower(t *testing.T) {
	num := writeableAsSumOfNthPower(5)
	if sum(num) != 443839 {
		t.Errorf("sum = %v, want %v", sum(num), 443839)
	}
}
