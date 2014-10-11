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

var (
	listFlag *bool
)

func main() {
	url := parse()
	open.Start(url)
}

func parse() string {
	args := parseFlags()
	engine, args := parseEngine(args)
	query := parseQuery(args)
	url := fmt.Sprintf(engine, query)
	return url
}

func parseFlags() []string {
	flag.Parse()
	args := flag.Args()
	if *listFlag {
		printListEngines()
	}
	if len(args) == 0 {
		flag.Usage()
	}
	return flag.Args()
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

func init() {
	const (
		help_usage = "print this help"
		list_usage = "list of defined engines"
	)
	flag.Bool("h", false, help_usage)
	listFlag = flag.Bool("list", false, list_usage)

	flag.Usage = func() {
		pattern := "search-me %s query string"
		engines := loadEngines()
		search_examples := ""
		for name := range engines {
			search_examples += "  " + fmt.Sprintf(pattern, name) + "\n"
		}
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Print(`
Usage:
  ` + fmt.Sprintf(pattern, "[engine]") + `

Examples:
` + search_examples)
		os.Exit(0)
	}
}

func printListEngines() {
	engines := loadEngines()
	configured_engines := ""
	for name, url := range engines {
		configured_engines += name + `: ` + url + "\n"
	}
	fmt.Print(`Configured engines:
` + configured_engines)
	os.Exit(0)
}
