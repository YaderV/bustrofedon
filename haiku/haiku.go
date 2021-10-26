package haiku

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

func CreateHaiku(jsonFile string) string {
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}

	haikus := []Verse{}

	json.Unmarshal([]byte(file), &haikus)

	rand.Seed(time.Now().UnixNano())
	max := len(haikus)
	haiku := make([]string, 3)

	firstVerse := strings.ToLower(haikus[rand.Intn(max)].A)
	middleVerse := strings.ToLower(haikus[rand.Intn(max)].B)
	lastVerse := strings.ToLower(haikus[rand.Intn(max)].C)

	haiku[0] = TransformFirstVerse(firstVerse)
	haiku[1] = CreateBoustrophedon(middleVerse)
	haiku[2] = TransformLastVerse(lastVerse)
	return strings.Join(haiku, "\n")
}

func TransformFirstVerse(verse string) string {
	verse = strings.ToUpper(verse[0:1]) + verse[1:]
	return CutLine(verse)
}

func TransformLastVerse(verse string) string {
	char := verse[len(verse)-1:]
	if char != "." {
		verse += "."
	}
	return verse
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
	newVerse = CutLine(strings.TrimSpace(newVerse))
	return newVerse
}

func CutLine(verse string) string {
	char := verse[len(verse)-1:]
	if char != "," && char != ":" {
		verse += ","
	}
	return verse

}
