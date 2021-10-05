package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Strip trailing whitespace and comments

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./example.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		s := scanner.Text()

		if !strings.HasPrefix(s, "#") {
			s = strings.TrimSpace(s)
			text = append(text, s)
		}
	}

	file.Close()

	wholeText := strings.Join(text, "\n")
	haikus := strings.Split(wholeText, "\n\n")
	haikus_length := len(haikus)
	fmt.Printf("%d Haikus Found\n", haikus_length)

	var haiku_output = make([]map[string]string, haikus_length)

	for index, text := range haikus {
		verses := strings.Split(text, "\n")

		var buckets [3]string

		for i, verse := range verses {
			verseWords := strings.Split(verse, " ")
			if len(verseWords) > 0 {
				buckets[i] = strings.Join(verseWords, " ")
			}
		}
		haiku_output[index] = map[string]string{
			"a": buckets[0],
			"b": buckets[1],
			"c": buckets[2],
		}
	}
	jsonString, err := json.MarshalIndent(haiku_output, "", " ")
	check(err)
	fmt.Printf("verses: %s\n", jsonString)
	_ = ioutil.WriteFile("verses.json", jsonString, 0644)
}
