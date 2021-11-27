package newpwd

import (
	"fmt"
	"testing"
)

func hasSymbols(s string, s2 string) bool {
	symbols := make(map[rune]bool)
	for _, c := range s2 {
		symbols[c] = true
	}

	for _, c := range s {
		if symbols[c] {
			return true
		}
	}
	return false
}

func TestMake(t *testing.T) {
	t.Run("no alphabet", func(t *testing.T) {
		p := Make(10, false, false, false, false)
		if p != "" {
			t.Errorf("expected \"\", got %q", p)
		}
	})

	for i := 1; i < 16; i++ {
		lower, upper, digits, punc := i&8 == 8, i&4 == 4, i&2 == 2, i&1 == 1
		for _, length := range []int{1, 5, 10} {
			t.Run(
				fmt.Sprintf("length=%d, lower=%v, upper=%v, digits=%v, punc=%v",
					length, lower, upper, digits, punc),
				func(t *testing.T) {
					pwd := Make(length, lower, upper, digits, punc)
					if len(pwd) != length {
						t.Errorf("expected length of %d, got %d", length, len(pwd))
					}

					if !lower {
						if hasSymbols(pwd, AsciiLower) {
							t.Errorf("expected string withous AsciiLower, got %s", pwd)
						}
					}

					if !upper {
						if hasSymbols(pwd, AsciiUpper) {
							t.Errorf("expected string withous AsciiUpper, got %s", pwd)
						}
					}

					if !digits {
						if hasSymbols(pwd, Digits) {
							t.Errorf("expected string withous Digits, got %s", pwd)
						}
					}

					if !punc {
						if hasSymbols(pwd, Punctuation) {
							t.Errorf("expected string withous Punctuation, got %s", pwd)
						}
					}

				})
		}
	}
}
