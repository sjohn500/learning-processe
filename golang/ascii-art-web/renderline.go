package main

import "strings"

func RenderLine(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	var builder strings.Builder

	lines := strings.Split(input, "\n")
	for _, word := range lines {
		if word == "" {
			builder.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range word {
				text, ok := banner[ch]
				if !ok {
					continue
				}
				builder.WriteString(text[row])
			}
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
