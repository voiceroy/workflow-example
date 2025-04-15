package main

import "testing"

func TestReturnHello(t *testing.T) {
	want := "Hello, World!"
	if got := ReturnHello(); got != want {
		t.Errorf("Wanted %s got %s", want, ReturnHello())
	}
}
