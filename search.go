package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
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
	engines := map[string]string{"main": "http://duckduckgo.com?q=%s"}
	loadEnginesFile(os.Getenv("HOME")+"/.search", &engines)
	return engines
}

func loadEnginesFile(path string, engines *map[string]string) {
	enginesPtr := *engines
	buffer, _ := ioutil.ReadFile(path)
	file := string(buffer)
	lines := strings.Split(file, "\n")
	for k := range lines {
		line := trim(lines[k])
		engine := RegSplit(line, `\s+`)
		if len(engine) == 2 {
			enginesPtr[engine[0]] = engine[1]
		}
	}
}

// Trim comments
func trim(str string) string {
	re := regexp.MustCompile(`(^|\s)#(.*)$`)
	line := strings.TrimSpace(str)
	return re.ReplaceAllLiteralString(line, "")
}

func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
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
