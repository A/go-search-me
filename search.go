package main

import (
	"flag"
	// "fmt"
	"io/ioutil"
	"net/url"
	// "os"
	"encoding/json"
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

// TODO: Support search with default engine
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

// TODO: load json from engines.json && ~/.search.json
func loadEngines() map[string]string {
	file, _ := ioutil.ReadFile("engines.json")
	content := []byte(string(file))
	engines := make(map[string]string)
	json.Unmarshal(content, &engines)
	return engines
}
