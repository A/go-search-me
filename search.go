package main

import (
	"flag"
	// "fmt"
	"net/url"
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

// TODO: load json from engines.json && ~/.search.json
func loadEngines() map[string]string {
	engines := make(map[string]string)
	engines["google"] = "http://google.com/#q=%s"
	engines["ddg"] = "http://duckduckgo.com/?q=%s"
	return engines
}
