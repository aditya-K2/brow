package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Query struct {
	Arr []interface{}
}

func NewQuery(q []interface{}) *Query {
	return &Query{
		q,
	}
}

func (q *Query) String() string { return fmt.Sprint(q.Arr) }

func (q *Query) Set(value string) error {
	q.Arr = make([]interface{}, 0)
	if value == "" || strings.ReplaceAll(value, ",", "") == "" {
		return errors.New("Empty Query!")
	}
	for _, s := range strings.Split(value, ",") {
		_s := strings.TrimSpace(s)
		if _s != "" {
			q.Arr = append(q.Arr, _s)
		}
	}
	return nil
}

var (
	fieldSepartor        = "|"
	subsystem            = "urls" // History
	query                = NewQuery([]interface{}{"*"})
	configDir, configErr = os.UserConfigDir()
	rhpath               = "/BraveSoftware/Brave-Browser/Default/History"
	path                 = configDir + rhpath
	update               = false
	seeAvailableOptions  = false
)

func ParseFlags() {
	flag.StringVar(&subsystem, "s", subsystem, "Specify Subsystem. Either h[istory]/d[ownloads]")
	flag.StringVar(&fieldSepartor, "f", fieldSepartor, "Specify the Field Separator which will be used in the Output.")
	flag.Var(query, "q", "Query for the current Subsystem")
	flag.BoolVar(&update, "u", update, "Update the Current Database")
	flag.BoolVar(&seeAvailableOptions, "a", seeAvailableOptions, "See All Available Options that can be used to query the Subsystem.")
	flag.StringVar(&path, "p", path, "Specify Custom Path for History")
	flag.Parse()
}

func main() {
	ParseFlags()
	if configErr != nil && path == rhpath {
		PrintC("RED", "Couldn't Get user's Config Directory!\n")
		panic(configErr)
	}
	if subsystem == "history" {
		subsystem = "urls"
	} else if subsystem != "downloads" && subsystem != "urls" {
		PrintC("RED", fmt.Sprintf("No Subystem with Name %s Found", subsystem)+"\n")
		os.Exit(-1)
	}
	ConnectDatabase()
	if seeAvailableOptions {
		GetSubsytemInfo()
		os.Exit(0)
	}
}
