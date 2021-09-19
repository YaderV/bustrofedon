package main

import (
	"fmt"
	"bufio"
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
	scanner.Split(bufio.ScanLines)
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
	fmt.Printf("%d Haikus Found", len(haikus))

}
