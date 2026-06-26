package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("could not read file")
	}

	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(lines) != 856 {
		return nil, errors.New("incomplete lines")
	}
	fontMap := make(map[rune][]string)
	ascii := rune(32)
	for i := 1; i < len(lines); i += 9 {
		if ascii < 126 {
			break
		}
		fontMap[ascii] = lines[i : i+8]
		ascii++
	}
	if len(fontMap) != 96 {
		return nil, errors.New("incomplete ascii charaters")
	}
	return fontMap, nil
}
