package main

import (
	"math/rand"
	"testing"
	"time"
)

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s) - 1 - i] {
			return false
		}
	}
	return true
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}

	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("ГПСЧ инициализирован: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}