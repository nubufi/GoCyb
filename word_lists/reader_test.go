package wordlists

import (
	"fmt"
	"testing"
)

func TestLoadWordList(t *testing.T) {
	wordList := LoadWordlist("Subdomain.txt")

	if len(wordList) == 0 {
		t.Errorf("No records found")
	}
}

func TestGetMainPath(t *testing.T) {
	path := getMainPath()

	fmt.Println(path)
}
