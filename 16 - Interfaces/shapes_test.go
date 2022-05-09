package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rec := Rectangle{10, 12}
		expected := float64(120)
		result := rec.Area()

		if result != expected {
			t.Fatalf("Resulting area %f differs from expeted area %f", result, expected)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}
		expected := float64(math.Pi * 100)

		result := circ.Area()

		if result != expected {
			t.Fatalf("Resulting area %f differs from expeted area %f", result, expected)
		}
	})
}
