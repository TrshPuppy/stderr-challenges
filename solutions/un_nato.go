package main

import (
	"fmt"
	"strings"
)

func main() {
	nato := "ALFA PAPA ROMEO INDIA LIMA  FOXTROT OSCAR OSCAR LIMA"
	decoded := []string{}
	for _, w := range strings.Split(nato, "  ") {
		for _, c := range strings.Split(w, " ") {
			decoded = append(decoded, string(c[0]))
		}
		decoded = append(decoded, " ")
	}

	fmt.Printf("Decoded: %s\n", string(strings.Join(decoded, "")))

}
