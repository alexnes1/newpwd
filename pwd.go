package newpwd

import (
	"math/rand"
)

const AsciiLower = "abcdefghijklmnopqrstuvwxyz"
const AsciiUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Digits = "0123456789"
const Punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

func Make(length int, useAsciiLower bool, useAsciiUpper bool, useDigits bool, usePunctuation bool) string {
	alphabet := []rune{}
	if useAsciiLower {
		alphabet = append(alphabet, []rune(AsciiLower)...)
	}
	if useAsciiUpper {
		alphabet = append(alphabet, []rune(AsciiUpper)...)
	}
	if useDigits {
		alphabet = append(alphabet, []rune(Digits)...)
	}
	if usePunctuation {
		alphabet = append(alphabet, []rune(Punctuation)...)
	}

	if len(alphabet) == 0 {
		return ""
	}

	pwd := []rune{}

	for i := 0; i < length; i++ {
		pos := rand.Int31n(int32(len(alphabet)))
		pwd = append(pwd, alphabet[pos])
	}

	return string(pwd)
}
