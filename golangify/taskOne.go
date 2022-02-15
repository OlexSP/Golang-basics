package main

import (
	"fmt"
	"strings"
)

// sends to downstream strings from the string slice
func sourceGopher(downstream chan string) {
	sourse := []string{"Vasil Terking", "Dara the dog", "Dara the dog", "Brother in low"}
	for _, v := range sourse {
		downstream <- v
	}
	close(downstream)
}

// filter dublicates and sends unique strings to the downstream
func dublicateFilter(upstream, downstream chan string) {
	itemPr := ""
	for item := range upstream {
		if item != itemPr {
			downstream <- item
			itemPr = item
		}
	}
	close(downstream)
}

// devides strings into words and sends words to the downstream
func wordDevider(upstream, downstream chan string) {
	for v := range upstream {
		for _, word := range strings.Fields(v) {
			downstream <- word
		}

	}
	close(downstream)
}

// prints strings from the channel
func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	c2 := make(chan string)

	go sourceGopher(c0)
	go dublicateFilter(c0, c1)
	go wordDevider(c1, c2)

	printGopher(c2)
}
