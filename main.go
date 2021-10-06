package main

import (
	"fmt"

	"github.com/yaderv/bustrofedon/corpus"
	"github.com/yaderv/bustrofedon/haiku"
)

// Strip trailing whitespace and comments

var INPUTFILE = "example.txt"
var OUTPUTFILE = "corpus.json"

func CreateJSON() {
	versesNumber, haikusString := corpus.CreateCorpusFromFile(INPUTFILE, OUTPUTFILE)
	fmt.Printf("verses: %d\n", versesNumber)
	fmt.Printf("%s Haikus Found\n", haikusString)

}

func main() {
	haiku := haiku.CreateHaiku(OUTPUTFILE)
	fmt.Println(haiku)
}
