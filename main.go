package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	HelpInfo = `
	brow -s [d(ownloads)/h(istory)] -q [query] -f [field separator] -p [path] -h -a

	-s		Specify Subsystem. Defaults to history.

	-h		Help

	-f		Specify the Field Separator which will be used in the Output, defaults to '|'

	-q		Query for the given substring

	-a		See All Available Options that can be used to query the Subsystem.

	-p		Specify Custom Path for History

	Note:
			The Letters in () are optional
			for e.g :
			history, h, his all will be considered history
	`
	fieldSepartor = "|"
	subsystem     = "history"
	query         = []string{"*"}
	configDir, _  = os.UserConfigDir()
	filePath      = configDir + "/BraveSoftware/Brave-Browser/Default/History"
)

func GetArgs() (f, s, p string, q []string) {
	f = fieldSepartor
	s = subsystem
	q = query
	p = filePath
	if len(os.Args) == 1 {
		PrintC("RED", "No Arguments Provided!\n")
		PrintC("GREEN", "Usage: ")
		fmt.Println(HelpInfo)
	} else {
		if os.Args[1] == "-h" {
			PrintC("GREEN", "Usage: ")
			fmt.Println(HelpInfo)
			os.Exit(0)
		}
		for k := range os.Args {
			if k != len(os.Args)-1 {
				if os.Args[k] == "-s" {
					if IsShortForm(os.Args[k+1], "downloads") {
						s = "downloads"
					} else if IsShortForm(os.Args[k+1], "history") {
						s = "history"
					} else {
						PrintC("RED", fmt.Sprintf("No Subystem with Name %s Found", os.Args[k+1])+"\n")
						os.Exit(-1)
					}
					k++
					continue
				} else if os.Args[k] == "-f" {
					f = os.Args[k+1]
					k++
					continue
				} else if os.Args[k] == "-q" {
					if (os.Args[k+1]) == "" || len(os.Args) == 2 {
						PrintC("RED", "Query Format Wrong\n")
						os.Exit(-1)
					} else {
						q = make([]string, 0)
						for _, s := range strings.Split(os.Args[k+1], ",") {
							q = append(q, strings.TrimSpace(s))
						}
					}
				} else if os.Args[k] == "-p" {
					if (os.Args[k+1]) == "" || len(os.Args) == 2 {
						PrintC("RED", "No Path for History Provided upon using -p flag\n")
						os.Exit(-1)
					} else {
						p = os.Args[k+1]
					}
				} else {
					if k != 0 {
						PrintC("RED", fmt.Sprintf("%s : No Such Flag Exists\n", os.Args[k]))
						os.Exit(-1)
					}
				}
			}
		}
	}
	return f, s, p, q
}

func main() {
	fieldSepartor, subsystem, filePath, query = GetArgs()
	fmt.Println(fieldSepartor, subsystem, filePath, query)
}
