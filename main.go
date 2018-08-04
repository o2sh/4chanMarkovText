package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//scrapp("pol")
	//scrapp("fit")
	//scrapp("b")
	//scrapp("biz")
	// Register command-line flags.
	numWords := flag.Int("words", 100, "maximum number of words to print")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")

	flag.Parse()                     // Parse command-line flags.
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

	c := NewChain(*prefixLen) // Initialize a new Chain.
	file, err := os.Open("./data/fit.txt")
	check(err)
	c.Build(bufio.NewReader(file)) // Build chains from standard input.
	text := c.Generate(*numWords)  // Generate text.
	fmt.Println(text)
}
