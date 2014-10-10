package main

import (
	"flag"
	"fmt"
	"net/url"
	// "os/exec"
	"strings"
)

func main() {
	url := parse()
	fmt.Println(url)

	// cmd := exec.Command("open", url)
	// cmd.Start()
}

func parse() string {
	args := parseFlags()
	engine := parseEngine(args)
	query := parseQuery(args)
	fmt.Println(query)
	return engine
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

// TODO: load json from engines.json && ~/.search.json
func loadEngines() map[string]string {
	engines := make(map[string]string)
	engines["google"] = "http://google.com/#q=%s"
	engines["ddg"] = "http://duckduckgo.com/?q=%s"
	return engines
}

// func getUrl(args []string) string {
// 	engine, query := parse(args)
// 	engine + query
// 	url := replace()
// 	return url
// }
//
// func parse(args []string) (string, []string) {
// 	query := QueryEscape(string())
// }
