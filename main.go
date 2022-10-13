package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Query struct {
	Arr []string
}

func NewQuery(q []string) *Query {
	return &Query{
		q,
	}
}

func (q *Query) String() string { return fmt.Sprint(q.Arr) }

func (q *Query) Set(value string) error {
	q.Arr = make([]string, 0)
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
	query                = NewQuery([]string{"*"})
	configDir, configErr = os.UserConfigDir()
	rhpath               = "/BraveSoftware/Brave-Browser/Default/History"
	opath                = configDir + rhpath
	cacheDir, cacheErr   = os.UserCacheDir()
	path                 = opath
	cachedPath           = cacheDir + "/BrowHistory"
	copy                 = false
	seeAvailableOptions  = false
	numberOfResults      = -1
	ascending            = false
)

func ParseFlags() {
	flag.StringVar(&subsystem, "s", subsystem, "Specify Subsystem. Either history(urls)/downloads")
	flag.StringVar(&fieldSepartor, "f", fieldSepartor, "Specify the Field Separator which will be used in the Output.")
	flag.Var(query, "q", "Query for the current Subsystem")
	flag.BoolVar(&copy, "c", copy, "Copy the Browser Database. Ignored if path is provided. Use this Flag in case Browser is Open as it locks the database.")
	flag.BoolVar(&seeAvailableOptions, "a", seeAvailableOptions, "See All Available Options that can be used to query the Subsystem.")
	flag.StringVar(&path, "p", path, "Specify Custom Path for History")
	flag.IntVar(&numberOfResults, "n", numberOfResults, "Specify Number of Results to be Displayed")
	flag.BoolVar(&ascending, "ascending", ascending, "Get Results in ascending order. Default is Descending Order")
	flag.Parse()
}

func main() {
	ParseFlags()
	if configErr != nil && path == rhpath {
		PrintC("RED", "Couldn't Get user's Config Directory!\n")
		panic(configErr)
	}
	if cacheErr != nil && path == opath {
		PrintC("RED", "Couldn't Get user's Cache Directory!\n")
		panic(cacheErr)
	}
	if copy && path == opath {
		if err := Copy(path, cachedPath); err != nil {
			PrintC("RED", fmt.Sprintf("Error Copying Database from %s to %s", path, cachedPath))
			panic(err)
		} else {
			path = cachedPath
		}
	}
	if subsystem == "history" {
		subsystem = "urls"
	} else if subsystem != "downloads" && subsystem != "urls" {
		PrintC("RED", fmt.Sprintf("No Subystem with Name %s Found", subsystem)+"\n")
		os.Exit(-1)
	}
	ConnectDatabase()
	if seeAvailableOptions {
		for k := range GetSubsytemInfo() {
			fmt.Println(k)
		}
		os.Exit(0)
	}
	for _, v := range GetData() {
		fmt.Println(v)
	}
}
