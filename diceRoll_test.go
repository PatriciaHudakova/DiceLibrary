package Library

import "testing"

// Rolls dice 20 times to establish some level of certainty
func TestRollDice(t *testing.T) {
	result := Roll()
	for i := 1; i == 20; i++ {
		if result > 6 || result < 0 {
			t.Errorf("the rolled result is %v", result)
		}
	}
}