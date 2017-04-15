package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	lines := make(map[string]int)

	if len(os.Args) < 2 {
		fmt.Println("usage: intersect [files]")
		os.Exit(1)
	}

	for _, path := range os.Args[1:] {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			count, ok := lines[line]

			if ok {
				lines[line] = count + 1
			} else {
				lines[line] = 1
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	intersect := make([]string, 0)

	all := len(os.Args) - 1

	for line, count := range lines {
		if count == all {
			intersect = append(intersect, line)
		}
	}

	sort.Strings(intersect)

	for _, line := range intersect {
		fmt.Println(line)
	}
}
