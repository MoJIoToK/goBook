// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"goBook/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
