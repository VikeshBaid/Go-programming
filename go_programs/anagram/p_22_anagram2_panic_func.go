package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

func sortRunes(str string) string {
	runes := bytes.Runes([]byte(str))
	var temp rune
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				temp = runes[i]
				runes[i], runes[j] = runes[j], temp
			}

		}
	}
	return string(runes)
}

func load(fname string) ([]string, error) {
	if fname == "" {
		return nil, errors.New("Dictionary file name cannot be emtpy.")
	}

	// attempt to load dict file.
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file %s: %s", fname, err)
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

// write function writes the map of anagram to a file file specified by name
// fname. The use of os.OpenFile( instead of os.Create ) is to demonstrate
// the use of panic() which is generated when OpenFile returns as error
func write(fname string, anagrams map[string][]string) {
	file, err := os.OpenFile(fname, os.O_WRONLY+os.O_CREATE+os.O_EXCL, 0644)
	if err != nil {
		msg := fmt.Sprintf("Unable to create output file: %v", err)
		panic(msg)
	}
	defer file.Close()

	for k, v := range anagrams {
		output := fmt.Sprintf("%s -> %v\n", k, v)
		file.WriteString(output)
	}
}

// mapWords maps each word to its associate signature
func mapWords(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		wordSig := sortRunes(word)
		anagrams[wordSig] = append(anagrams[wordSig], word)
	}

	return anagrams
}

func main() {
	words, err := load("dict.txt")
	if err != nil {
		fmt.Println("unable to load file:", err)
		os.Exit(1)
	}

	//generate map of anagrams
	anagrams := mapWords(words)

	// write output in file
	write("out.txt", anagrams)
}
