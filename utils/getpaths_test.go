package utils

import (
"testing"
)

func TestContains(t *testing .T){
	input := []string{"A", "B", "C", "D"}
	expected := true
	got := Contains(input, "B")

	if expected != got{
		t.Errorf("Test Contains failed got %v, expected %v", got, expected)
	}
}
