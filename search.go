package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

func main() {
	url := parse()
	open.Run(url)
}

func parse() string {
	args := parseFlags()
	if len(args) == 0 {
		usage()
	}
	engine, args := parseEngine(args)
	query := parseQuery(args)
	url := strings.Replace(engine, "%s", query, -1)
	return url
}

func parseFlags() []string {
	flag.Parse()
	return flag.Args()
}

func usage() {
	pattern := "search-me %s query string"
	engines := loadEngines()
	search_examples := ""
	for name := range engines {
		search_examples += fmt.Sprintf(pattern, name) + "\n"
	}
	fmt.Printf(`Usage:

Pattern:

` + fmt.Sprintf(pattern, "[engine]") + `

Examples:

` + search_examples)
	os.Exit(1)
}

func parseEngine(args []string) (string, []string) {
	engines := loadEngines()
	name := args[0]
	if engine, ok := engines[name]; ok {
		args := args[1:]
		return engine, args
	}
	engine := engines["main"]
	return engine, args
}

func parseQuery(args []string) string {
	query := strings.Join(args, " ")
	return url.QueryEscape(query)
}

func loadEngines() map[string]string {
	engines := make(map[string]string)
	loadEnginesFile("engines.json", &engines)
	loadEnginesFile(os.Getenv("HOME")+"/.search.json", &engines)
	return engines
}

func loadEnginesFile(path string, engines *map[string]string) {
	buffer, _ := ioutil.ReadFile(path)
	json.Unmarshal(buffer, &engines)
}
