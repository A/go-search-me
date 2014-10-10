package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	url := parse()
	cmd := exec.Command("open", url)
	cmd.Start()
}

func parse() string {
	args := parseFlags()
	engine := parseEngine(args)
	query := parseQuery(args)
	url := strings.Replace(engine, "%s", query, -1)
	return url
}

func parseFlags() []string {
	flag.Parse()
	return flag.Args()
}

func parseEngine(args []string) string {
	engines := loadEngines()
	name := args[0]
	engine := engines[name]
	return engine
}

func parseQuery(args []string) string {
	query := strings.Join(args, " ")
	return url.QueryEscape(query)
}

type Engine struct {
	url string
}

func loadEngines() map[string]string {
	engines := make(map[string]string)
	loadEnginesFile("engines.json", &engines)
	loadEnginesFile(os.Getenv("HOME")+"/.search.json", &engines)
	return engines
}

func loadEnginesFile(path string, engines *map[string]string) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(buffer, &engines)
}
