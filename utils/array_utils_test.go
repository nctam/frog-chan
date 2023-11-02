package utils

import "testing"

func TestStringContains(t *testing.T) {
	expected := true
	input := []string{"frog-chan", "dill-chan", "xxxx"}
	output := StringContains(input, "frog-chan")
	if output != expected {
		t.Errorf("Got: %v, output: %v\n", output, expected)
	}
}

func TestStringNotContains(t *testing.T) {
	expected := true
	input := []string{"frog-chan", "dill-chan", "xxxx"}
	output := StringContains(input, "hello")
	if output == expected {
		t.Errorf("Got: %v, output: %v\n", output, expected)
	}
}
