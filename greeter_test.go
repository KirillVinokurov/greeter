package greeter

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Viktor"
	want := regexp.MustCompile("Hello " + name)
	msg, err := Greet(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet(name) = %q, want match for %#q, nil`, msg, want)
	}
}
func TestWithoutName(t *testing.T) {
	_, err := Greet("")
	if err == nil {
		t.Fatal("emty name must cause an error")
	}
}
