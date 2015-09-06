package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tj/docopt"
)

const (
	version = "1.0.0"
	usage   = `DiLaurentis.

DiLaurentis pretty prints JSON documents from standard input.

Usage:
  dilaurentis [--indentation indent]
  dilaurentis -h | --help
  dilaurentis --version

Options:
  --indentation indent    Indentation Level [default: 2]
  -h --help               Show this screen.
  --version               Show version.`
)

func main() {
	run(nil, os.Stdin, os.Stdout)
}

func run(argv []string, in io.Reader, out io.Writer) {
	arguments, err := docopt.Parse(usage, argv, true, version, false)
	check(err)

	indent, err := strconv.Atoi(arguments["--indentation"].(string))
	check(err)

	DiLaurentis(in, out, indent)
}

func DiLaurentis(in io.Reader, out io.Writer, indent int) {
	var data interface{}

	err := json.NewDecoder(in).Decode(&data)
	check(err)

	o, err := json.MarshalIndent(data, "", strings.Repeat(" ", indent))
	check(err)

	out.Write(o)
	fmt.Fprintln(out)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
