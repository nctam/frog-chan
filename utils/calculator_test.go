package utils

import (
	"fmt"
	"golang.org/x/exp/rand"
	"testing"
	"time"
)

func TestShouldTarget(t *testing.T) {
	t.Run("should target", func(t *testing.T) {
		counter := 0
		for i := 0; i < 100; i++ {
			rand.Seed(uint64(time.Now().UnixNano()))
			randValue := rand.Int()
			if randValue%5 == 0 {
				counter += 1
			}
		}
		fmt.Println(counter)
	})
}
