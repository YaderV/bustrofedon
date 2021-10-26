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
	lastVerse := strings.ToLower(haikus[rand.Intn(max)].C)

	haiku[0] = TransformFirstVerse(firstVerse)
	haiku[1] = strings.ToLower(haikus[rand.Intn(max)].B)
	haiku[2] = TransformLastVerse(lastVerse)
	return strings.Join(haiku, "\n")
}

func TransformFirstVerse(verse string) string {
	return strings.ToUpper(verse[0:1]) + verse[1:]
}

func TransformLastVerse(verse string) string {
	char := verse[len(verse)-1:]
	if char != "." {
		verse += "."
	}
	return verse
}
