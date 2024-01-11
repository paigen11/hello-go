package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with name, checking for valid return value
func TestHelloName(t *testing.T) {
	name := "May"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("May")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("May") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with empty string, checking for error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
