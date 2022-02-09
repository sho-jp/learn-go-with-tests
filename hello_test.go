package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Chris")
  expect := "Hello, Chris"

  if got != expect {
    t.Errorf("%q != %q", got, expect)
  }
}
