package iteration

import "strings"

const repeatCount = 5

// Output string iteration.
func Repeat(char string, count int) (repeated string) {
	return strings.Repeat(char, count)
}
