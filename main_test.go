package main

import "testing"

func TestReturnHello(t *testing.T) {
	want := "Hello, World!"
	if got := ReturnHello(); got != want {
		t.Errorf("Wanted %s got %s", want, ReturnHello())
	}
}

func TestReturnInteger(t *testing.T) {
	got := ReturnInteger()
	if got%2 == 0 {
		t.Logf("Integer is even: %d\n", got)
	} else {
		t.Logf("Integer is odd: %d\n", got)
	}
}
