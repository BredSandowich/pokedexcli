package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	cleanedInput := strings.Fields(lowered)
	
	return cleanedInput
}