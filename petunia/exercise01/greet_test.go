package greet

import "testing"

func TestSayHello_Ingrid(t *testing.T) {
	name := "Ingrid"
	want := "Hello, Coxis"
	if got := SayHello(name); got != want {
		t.Errorf("SayHello() = %q, want %q", got, want)
	}
}

func TestSayHello_SomeoneElse(t *testing.T) {
	name := "Darragh"
	want := "Hello, " + name
	if got := SayHello(name); got != want {
		t.Errorf("SayHello() = %q, want %q", got, want)
	}
}
