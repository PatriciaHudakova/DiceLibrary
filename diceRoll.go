package Library

import (
	"math/rand"
	"time"
)

func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	return rand.Intn(max - min + 1) + min
}