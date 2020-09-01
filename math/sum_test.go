package math_test

import (
	"testing"

	"github.com/xabi93/cloud-function/math"
)

func TestSum(t *testing.T) {
	a, b, expected := 1, 1, 2

	if result := math.Sum(a, b); result != expected {
		t.Errorf("the sum of %d and %d expected to be %d, but it was %d", a, b, expected, result)
	}
}
