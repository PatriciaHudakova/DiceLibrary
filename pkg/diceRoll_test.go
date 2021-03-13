package pkg

import "testing"

func TestRollDice(t *testing.T) {
	result := rollDice()
	if result > 6 || result == 0 {
		t.Errorf("the rolled result is %v", result)
	}
}