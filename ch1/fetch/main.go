// Fetch выводит ответ на запрос по заданному URL.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const filename = "test.txt"

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение1 %s: %v\n", url, err)
			os.Exit(1)
		}

		file, err := os.Create("out.txt")
		if err != nil {
			log.Println("Open")
		}

		writer := bufio.NewWriter(file)

		_, err = io.Copy(writer, resp.Body)
		writer.Flush()
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение2 %s: %v\n", url, err)
			os.Exit(1)
		}
		defer file.Close()
		//writer.Flush()
		//fmt.Printf("%s", b)
	}

}
