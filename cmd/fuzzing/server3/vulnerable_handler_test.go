// vulnerable_handler_test.go
package main

import (
	"strings"
	"testing"
)

func FuzzVulnerableHandler(f *testing.F) {
	f.Add("fuzzing_input")
	f.Fuzz(func(t *testing.T, input string) {
		buffer := make([]byte, 1024)
		copy(buffer, input)

		if strings.Contains(string(buffer), "abc123") {
			t.Error("¡Vulnerabilidad encontrada!")
		}
	})
}
