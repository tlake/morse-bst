package main

import (
	"bufio"
	"fmt"
	"github.com/tlake/morse-bst/bst"
	"log"
	"os"
	"strings"
)

func removeEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}

func main() {
	tree, err := bst.NewMorseTree()
	if err != nil {
		log.Fatal("could not initialize tree")
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Separate characters with '/' and words with ' '.")
	fmt.Printf("\nEnter morse code\n> ")

	for scanner.Scan() {
		text := scanner.Text()
		words := removeEmpty(strings.Split(text, " "))

		var results []string

		for _, word := range words {
			var result string
			chars := removeEmpty(strings.Split(word, "/"))

			fmt.Printf("chars: %v\n", chars)

			for _, morse := range chars {
				result += tree.Decode(morse)
			}

			results = append(results, result)
		}

		fmt.Println(strings.Join(results, " "))
		fmt.Printf("\nEnter morse code\n> ")
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
	}
}
