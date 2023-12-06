package utils

import (
	"fmt"
	"testing"
)

func TestMakeDecision(t *testing.T) {
	t.Run("make decision", func(t *testing.T) {
		counter, decision := makeDecision(0.2, 100)
		fmt.Printf("Total votes: %d ==> decision: %v", counter, decision)
	})
}
