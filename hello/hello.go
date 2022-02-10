package main

import "fmt"

func main() {
  fmt.Println(Hello("Chris"))
}

func Hello(name string) string {
  return fmt.Sprintf("Hello, %s", name)
}
