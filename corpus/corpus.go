package corpus

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateCorpusFromFile(inputFile string, outputFile string) (int, []byte) {
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
	haikusLength := len(haikus)

	var haiku_output = make([]map[string]string, haikusLength)

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
	_ = ioutil.WriteFile(outputFile, jsonString, 0644)
	return haikusLength, jsonString
}
