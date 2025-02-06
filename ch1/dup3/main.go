package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		space := []byte("\n")
		data = append(data, space...)
		s := strings.Split(string(data), "\n")

		for _, line := range s {
			fmt.Println(line)
			count[line]++
		}
	}
	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
