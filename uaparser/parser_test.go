package uaparser

import (
	"testing"
)

func testAllMatchesReplacement(t *testing.T, expected, pattern string, values []string) {
	actual := allMatchesReplacement(pattern, values)
	if actual != expected {
		t.Fatalf("Expected '%s', but got '%s' while replacing tokens in '%s' with %v", expected, actual, pattern, values)
	}
}

func TestAllMatchesReplacementEmptyPattern(t *testing.T) {
	testAllMatchesReplacement(t, "", "", []string{"whole match"})
	testAllMatchesReplacement(t, "", "", []string{"whole match", "a"})
}

func TestAllMatchesReplacementNoTokensInPattern(t *testing.T) {
	testAllMatchesReplacement(t, "foo", "foo", []string{"whole match"})
	testAllMatchesReplacement(t, "foo", "foo", []string{"whole match", "a"})
	testAllMatchesReplacement(t, "100$", "100$", []string{"whole match", "a"})
	testAllMatchesReplacement(t, "$not-a-token", "$not-a-token", []string{"whole match", "a"})
	testAllMatchesReplacement(t, "$ ", "$ ", []string{"whole match", "a"})
	testAllMatchesReplacement(t, "$$$", "$$$", []string{"whole match", "a"})
}

func TestAllMatchesReplacementNoValueForToken(t *testing.T) {
	testAllMatchesReplacement(t, "$1 $3 $2", "$1 $3 $2", []string{"whole match"})
	testAllMatchesReplacement(t, "one $3 $2", "$1 $3 $2", []string{"whole match", "one"})
	testAllMatchesReplacement(t, "one $3 two", "$1 $3 $2", []string{"whole match", "one", "two"})
}

func TestAllMatchesReplacementMultidigitToken(t *testing.T) {
	testAllMatchesReplacement(t, "$10 > $1", "$10 > $1", []string{"whole match"})
	testAllMatchesReplacement(t, "$10 > one", "$10 > $1", []string{"whole match", "one"})
	testAllMatchesReplacement(t, "$10 > one", "$10 > $1", []string{"whole match", "one"})
	testAllMatchesReplacement(t, "ten > one", "$10 > $1",
		[]string{"whole match", "one", "2", "3", "4", "5", "6", "7", "8", "9", "ten"})
}

func TestAllMatchesReplacementTokenInsideToken(t *testing.T) {
	testAllMatchesReplacement(t, "one $1 $3 three", "$1 $2 $3", []string{"whole match", "one", "$1 $3", "three"})
}

func TestAllMatchesReplacementDuplicateToken(t *testing.T) {
	testAllMatchesReplacement(t, "a, a and a", "$1, $1 and $1", []string{"whole match", "a"})
}

func TestAllMatchesReplacementSkipsNotTokens(t *testing.T) {
	testAllMatchesReplacement(t, "$100", "$$1", []string{"whole match", "100"})
	testAllMatchesReplacement(t, "100$", "$1$", []string{"whole match", "100"})
	testAllMatchesReplacement(t, "$0", "$0", []string{"whole match", "100"})
	testAllMatchesReplacement(t, "$+1", "$+1", []string{"whole match", "100"})
	testAllMatchesReplacement(t, "$-1", "$-1", []string{"whole match", "100"})
}

func TestAllMatchesReplacementTooBigIndex(t *testing.T) {
	testAllMatchesReplacement(t, "$99999999999999999999", "$99999999999999999999", []string{"whole match"})
}
