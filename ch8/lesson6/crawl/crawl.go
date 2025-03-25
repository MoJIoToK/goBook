// Crawl1 crawls web links starting with the command-line arguments.
//
// This version quickly exhausts available file descriptors
// due to excessive concurrent calls to links.Extract.
//
// Also, it never terminates because the worklist is never closed.
package main

import (
	"fmt"
	"log"
	"os"

	"goBook/ch5/lesson6/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	// Запуск с аргументами командной строки
	go func() { worklist <- os.Args[1:] }()

	// Параллельное сканирование
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
