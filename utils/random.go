package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generated a random number between min and max
func RandomInt(min, max, int64) int64 {
	return min + rand.Int63n(max-min+1)
}
