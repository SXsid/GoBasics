package iteration

import "strings"

func Repeat(word string, time int) string {
	var value strings.Builder
	for i := 0; i < time; i++ {
		value.WriteString(word)
	}
	return value.String()
}
