package utils

import (
	"fmt"
	"testing"
)

func TestExtractMessage(t *testing.T) {
	expected := false
	msgs := []string{"nwng", "bot lor chửi ech chan", "chửi", "nứng"}
	for _, msg := range msgs {
		output := ExtractMessage("chửi", msg)
		if output != expected {
			t.Errorf("Got: %v, expect: %v -> Got is not matched expect\n", output, expected)
		} else {
			fmt.Printf("Got: %v, expect: %v -> Got is matched expect\n", output, expected)
		}
	}
}
