package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

// sort letters in a words (i.e "morning" -> "gimnnor")
func sortRunes(str string) string {
	runes := bytes.Runes([]byte(str))
	var temp rune
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				temp = runes[i]
				runes[i] = runes[j]
				runes[j] = temp
			}
		}
	}
	return string(runes)
}

// load loads a the content of the specified file's name into memory
// as an slice (array) of strings.
// function will return error if there will be a probelm in loading file
// scanner.Split() uses a function that knows to split the file records.
// here provided function bufio.ScanLines will do the job of splitting the
// records from the file and return a word for each line read and
// saves in slice of words
func load(fname string) ([]string, error) {
	if fname == "" {
		return nil, errors.New("Dictionary file name cannot be empty")
	}

	// attempt to load file
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// loads dictionary words in memory
	words, err := load("dict.txt")

	//test err for non-nil value
	//if err != nil, then the error is handled.
	// In this context, the program ends.
	if err != nil {
		fmt.Println("Unable to load file:", err)
		os.Exit(1)
	}

	//
	anagrams := make(map[string][]string)
	for _, word := range words {
		wordSig := sortRunes(word)
		anagrams[wordSig] = append(anagrams[wordSig], word)
	}

	// prints the result
	for k, v := range anagrams {
		fmt.Println(k, "->", v)
	}
}
