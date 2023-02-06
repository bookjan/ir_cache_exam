package main

import (
	"bufio"
	"fmt"
	"os"
)

var cachedTermMap = make(map[string]int)

func main() {
	fmt.Println(`Go's information retrieval exam.`)
	fmt.Println(`
commands: 
  exit() - exit the program
  list() - list all the “being used” terms 
	`)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n# ") // Prompt

		scanner.Scan()
		text := scanner.Text()

		fmt.Println(text)

		if text == "exit()" {
			break
		}

		if text == "list()" {
			for k, v := range cachedTermMap {
				if v >= 2 {
					fmt.Printf("term: %v, hits: %v\n", k, v)
				}
			}
		}

		if text != "" {
			hits, ok := cachedTermMap[text]
			if ok {
				hits++
				cachedTermMap[text] = hits
			} else {
				cachedTermMap[text] = 1
			}
		}

		// Go to get result by term from data source...
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}
