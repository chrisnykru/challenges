package main

import (
	"os"
	"testing"
)

func TestNameScore(t *testing.T) {
	score := nameScore("COLIN")
	if score != 53 {
		t.Errorf("score = %v, want %v", score, 53)
	}
}

func Test(t *testing.T) {
	f, err := os.Open("names.txt")
	if err != nil {
		t.Fatal(err)
	}

	totalScore := totalNameScores(f)
	expected := int64(871198282)
	if totalScore != expected {
		t.Errorf("total score = %v, want %v", totalScore, expected)
	}
}
