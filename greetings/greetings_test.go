package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Neto"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Neto")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Neto") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
