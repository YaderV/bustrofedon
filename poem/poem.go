package poem

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Verse struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

const HAIKULEN = 3
const QUARTETLET = 4

func LoadSource(jsonFile string) []Verse {
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	source := []Verse{}
	json.Unmarshal([]byte(file), &source)
	return source
}

func CreateHaiku(jsonFile string) string {
	source := LoadSource(jsonFile)
	rand.Seed(time.Now().UnixNano())
	max := len(source)
	poem := make([]string, HAIKULEN)

	firstVerse := strings.ToLower(source[rand.Intn(max)].A)
	middleVerse := strings.ToLower(source[rand.Intn(max)].B)
	lastVerse := strings.ToLower(source[rand.Intn(max)].C)

	poem[0] = TransformFirstVerse(firstVerse)
	poem[1] = CutVerse(CreateBoustrophedon(middleVerse))
	poem[2] = TransformLastVerse(lastVerse)
	return strings.Join(poem, "\n")
}

func CreateQuartet(jsonFile string) string {
	source := LoadSource(jsonFile)
	rand.Seed(time.Now().UnixNano())
	max := len(source)
	poem := make([]string, QUARTETLET)

	firstVerse := strings.ToLower(source[rand.Intn(max)].A)
	secondVerse := strings.ToLower(source[rand.Intn(max)].B)
	thirdVerse := strings.ToLower(source[rand.Intn(max)].A)
	lastVerse := strings.ToLower(source[rand.Intn(max)].C)

	poem[0] = TransformFirstVerse(firstVerse)
	poem[1] = CutVerse(CreateBoustrophedon(secondVerse))
	poem[2] = CutVerse(thirdVerse)
	poem[3] = TransformLastVerse(CreateBoustrophedon(lastVerse))
	return strings.Join(poem, "\n")

}

func TransformFirstVerse(verse string) string {
	verse = strings.ToUpper(verse[0:1]) + verse[1:]
	return CutVerse(verse)
}

func TransformLastVerse(verse string) string {
	char := verse[len(verse)-1:]
	if char != "." {
		verse += "."
	}
	return strings.TrimSpace(verse)
}

func CreateBoustrophedon(verse string) string {
	words := strings.Split(verse, " ")
	invertedWords := make([]string, len(words))
	index := 0
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) > 3 {
			invertedWords[index] = words[i]
			index += 1
		}
	}
	newVerse := strings.Join(invertedWords, " ")
	return strings.TrimSpace(newVerse)
}

func CutVerse(verse string) string {
	char := verse[len(verse)-1:]
	if char != "," && char != ":" {
		verse += ","
	}
	return strings.TrimSpace(verse)

}
