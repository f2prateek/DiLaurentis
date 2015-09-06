package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

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
  --indentation indent    Indentation [default:   ]
  -h --help               Show this screen.
  --version               Show version.`
)

func main() {
	run(nil, os.Stdin, os.Stdout)
}

func run(argv []string, in io.Reader, out io.Writer) {
	arguments, err := docopt.Parse(usage, argv, true, version, false)
	check(err)

	indent := arguments["--indentation"].(string)
	DiLaurentis(in, out, indent)
}

func DiLaurentis(in io.Reader, out io.Writer, indent string) {
	var data interface{}

	err := json.NewDecoder(in).Decode(&data)
	check(err)

	o, err := json.MarshalIndent(data, "", indent)
	check(err)

	out.Write(o)
	fmt.Fprintln(out)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
