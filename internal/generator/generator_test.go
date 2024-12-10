package generator

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	gen := New()

	// Test GenerateInt
	t.Run("GenerateInt", func(t *testing.T) {
		min, max := 1, 10
		for i := 0; i < 100; i++ {
			val := gen.GenerateInt(min, max)
			if val < min || val > max {
				t.Errorf("GenerateInt(%d, %d) = %d; want value between %d and %d", 
					min, max, val, min, max)
			}
		}
	})

	// Test GenerateFloat
	t.Run("GenerateFloat", func(t *testing.T) {
		min, max := 0.0, 1.0
		for i := 0; i < 100; i++ {
			val := gen.GenerateFloat(min, max)
			if val < min || val > max {
				t.Errorf("GenerateFloat(%f, %f) = %f; want value between %f and %f", 
					min, max, val, min, max)
			}
		}
	})

	// Test GenerateString
	t.Run("GenerateString", func(t *testing.T) {
		length := 10
		str := gen.GenerateString(length)
		if len(str) != length {
			t.Errorf("GenerateString(%d) returned string of length %d; want %d", 
				length, len(str), length)
		}
	})

	// Test GenerateBool
	t.Run("GenerateBool", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			val := gen.GenerateBool()
			if val != true && val != false {
				t.Error("GenerateBool() returned non-boolean value")
			}
		}
	})
}