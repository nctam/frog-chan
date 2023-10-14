package utils

import (
	"context"
	"testing"
)

func TestShouldTarget(t *testing.T) {
	t.Run("should target", func(t *testing.T) {
		prob := 0.2
		set := 1000
		_ = ShouldReply(context.TODO(), prob, set)
	})
}
