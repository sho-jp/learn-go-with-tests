package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello()
  expect := "Hello Wolrd!"

  if got != expect {
    t.Errorf("%s != %s", got, expect)
  }

}
