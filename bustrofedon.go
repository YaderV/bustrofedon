package main

import (
	"flag"
	"fmt"

	"github.com/yaderv/bustrofedon/corpus"
	"github.com/yaderv/bustrofedon/haiku"
)

// Strip trailing whitespace and comments

var OUTPUTFILE = "corpus.json"

func CreateJSON(inputFile string) {
	versesNumber, haikusString := corpus.CreateCorpusFromFile(inputFile, OUTPUTFILE)
	fmt.Printf("verses: %d\n", versesNumber)
	fmt.Printf("%s Haikus Found\n", haikusString)

}

func main() {
	source := flag.String("source", "", "The input file to build the haikus source")
	create := flag.Bool("haiku", false, "Create a haiku based on the haikus source")
	flag.Parse()

	if *source != "" {
		fmt.Printf("Reading from: %s", *source)
		CreateJSON(*source)
	}

	if *create {
		haiku := haiku.CreateHaiku(OUTPUTFILE)
		fmt.Println(haiku)
	}

}
