package wordlists

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadWordlist loads the given wordlist from word_lists directory.
// It returns a slice of strings containing the words in the wordlist.
func LoadWordlist(fileName string) []string {
	mainPath := getMainPath()
	filePath := fmt.Sprintf("%s/word_lists/%s", mainPath, fileName)

	f, err := os.Open(filePath)
	if err != nil {
		return []string{}
	}

	defer f.Close()

	var wordList []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	return wordList
}

// getMainPath returns the path up to the GoCyb directory.
func getMainPath() string {
	callerPath, _ := os.Getwd()

	targetDir := "GoCyb"

	index := strings.Index(callerPath, targetDir)
	pathUpToTarget := callerPath[:index+len(targetDir)]

	return pathUpToTarget
}
