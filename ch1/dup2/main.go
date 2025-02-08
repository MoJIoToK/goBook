// Dup2 выводит текст каждой строки, которая появляется во входных данных более одного раза.
// Программа читает стандартный ввод или список именованных файлов.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	//Стандартный ввод
	if len(files) == 0 {
		countLines(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, count)
			f.Close()
		}
	}
	for line, n := range count {
		fmt.Printf("%d\t%s\n", n, line)
	}

}

func countLines(f *os.File, count map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count[scanner.Text()]++
	}
}
