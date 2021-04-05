package library

import (
	"math/rand"
	"time"
)

func Roll() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	return rand.Intn(max - min + 1) + min
}