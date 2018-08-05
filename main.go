package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	//scrapp("pol")
	//scrapp("fit")
	scrapp("b")
	//scrapp("biz")

	log.SetFlags(0)
	log.SetPrefix("markov: ")
	input := flag.String("in", "./data/b.txt", "input file")
	n := flag.Int("n", 2, "number of words to use as prefix")
	wordsPerRun := flag.Int("words", 100, "number of words per run")
	startOnCapital := flag.Bool("capital", true, "start output with a capitalized prefix")
	stopAtSentence := flag.Bool("sentence", false, "end output at a sentence ending punctuation mark (after n words)")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	m, err := NewMarkovFromFile(*input, *n)
	check(err)

	err = m.Output(os.Stdout, *wordsPerRun, *startOnCapital, *stopAtSentence)
	check(err)
}
