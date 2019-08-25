package combos

import (
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	t.Run("single digit", func(t *testing.T) {
		got, _ := FromString("2")
		want := []string{"A", "B", "C"}
		assertCombos(t, got, want)
	})

	t.Run("digit with more than 3 characters combo", func(t *testing.T) {
		got, _ := FromString("7")
		want := []string{"P", "Q", "R", "S"}
		assertCombos(t, got, want)
	})

	t.Run("twice the same digit", func(t *testing.T) {
		got, _ := FromString("22")
		want := []string{
			"AA", "BA", "CA",
			"AB", "BB", "CB",
			"AC", "BC", "CC",
		}
		assertCombos(t, got, want)
	})

	t.Run("43556 can be translated as HELLO", func(t *testing.T) {
		input := "43556"
		got, _ := FromString(input)
		want := "HELLO"
		found := false

		for _, word := range got {
			if word == want {
				found = true
			}
		}

		if !found {
			t.Errorf("expected to find %q from %q", want, input)
		}
	})

	t.Run("non-digit input returns an error", func(t *testing.T) {
		got, err := FromString("H3LL0")

		if err == nil {
			t.Error("expected an error")
		}

		if got != nil {
			t.Errorf("Expected nil, got %v", got)
		}
	})
}

func BenchmarkFromString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FromString("43556")
	}
}

func assertCombos(t *testing.T, got, want []string) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
