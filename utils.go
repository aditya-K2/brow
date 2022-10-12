package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

func Copy(src, destination string) error {
	source, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	} else {
		err = ioutil.WriteFile(destination, source, 0644)
		if err != nil {
			return err
		}
		return nil
	}
}

func ExpandHomeDir(path string) string {
	HOME_DIR, _ := os.UserHomeDir()
	if strings.HasPrefix(path, "~/") {
		return filepath.Join(HOME_DIR, path[1:])
	} else if path == "~" {
		return HOME_DIR
	} else {
		return path
	}
}
