package main

import (
	"flag"
	"fmt"

	"github.com/yaderv/bustrofedon/corpus"
	"github.com/yaderv/bustrofedon/poem"
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
	poemType := flag.String("type", "haiku", "Poem type: haiku | quartet")
	flag.Parse()

	if *source != "" {
		fmt.Printf("Reading from: %s", *source)
		CreateJSON(*source)
	}

	if *poemType == "haiku" {
		haiku := poem.CreateHaiku(OUTPUTFILE)
		fmt.Println(haiku)
	} else if *poemType == "quartet" {
		haiku := poem.CreateQuartet(OUTPUTFILE)
		fmt.Println(haiku)
	} else {
		fmt.Println("Type option should be haiku or quartet")
	}

}
