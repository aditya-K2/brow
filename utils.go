package main

import "fmt"

var (
	COLORS map[string]string = map[string]string{
		"RESET":  "\033[0m",
		"RED":    "\033[31m",
		"GREEN":  "\033[32m",
		"YELLOW": "\033[33m",
		"BLUE":   "\033[34m",
		"PURPLE": "\033[35m",
		"CYAN":   "\033[36m",
		"GRAY":   "\033[37m",
		"WHITE":  "\033[97m",
	}
)

func PrintC(color, text string) {
	fmt.Print(COLORS[color] + text + COLORS["RESET"])
}

func IsShortForm(s string, o string) bool {
	if len(s) == 0 || (len(s) > len(o)) {
		return false
	} else {
		return s == o[:len(s)]
	}
}
