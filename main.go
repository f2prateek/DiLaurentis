package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	DiLaurentis(os.Stdin, os.Stdout)
}

func DiLaurentis(in io.Reader, out io.Writer) {
	var data interface{}

	err := json.NewDecoder(in).Decode(&data)
	check(err)

	o, err := json.MarshalIndent(data, "", "  ")
	check(err)

	out.Write(o)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
