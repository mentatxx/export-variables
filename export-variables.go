package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var outputType string = ""
var fileName string = ""

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseArguments(fileName *string, outputType *string) {
	flag.StringVar(outputType, "type", "json", "Output type: json (default), assignment")
	flag.StringVar(fileName, "file", "", "File name: Default empty - console")
	flag.Parse()
}

func printToFile(fileName string, content string) {
	if fileName != "" {
		// Write to file
		bytes := []byte(content)
		err := ioutil.WriteFile(fileName, bytes, 0644)
		check(err)
	} else {
		fmt.Println(content)
	}
}

func environmentToAssignment() string {
	env := os.Environ()
	content := strings.Join(env, "\n")
	return content
}

func environmentToJSON() string {
	env := os.Environ()
	envMap := make(map[string]string)
	for _, e := range env {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] != "" {
			envMap[pair[0]] = pair[1]
		}
	}
	content, err := json.Marshal(envMap)
	check(err)
	return string(content)
}

func main() {
	parseArguments(&fileName, &outputType)
	var content string = ""
	switch outputType {
	case "assignment":
		content = environmentToAssignment()
	case "json":
		content = environmentToJSON()
	default:
		panic("Invalid output type")
	}
	printToFile(fileName, content)
}
