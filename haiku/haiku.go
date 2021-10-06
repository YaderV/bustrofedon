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
	haiku[0] = haikus[rand.Intn(max)].A
	haiku[1] = haikus[rand.Intn(max)].B
	haiku[2] = haikus[rand.Intn(max)].C
	return strings.Join(haiku, "\n")
}
