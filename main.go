package main

import (
	"fmt"

	"github.com/yaderv/bustrofedon/corpus"
)

// Strip trailing whitespace and comments

var INPUTFILE = "example.txt"
var OUTPUTFILE = "corpus.json"

func main() {
	versesNumber, haikusString := corpus.CreateCorpusFromFile(INPUTFILE, OUTPUTFILE)
	fmt.Printf("verses: %d\n", versesNumber)
	fmt.Printf("%s Haikus Found\n", haikusString)
}
