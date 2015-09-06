package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDiLaurentis(t *testing.T) {
	cases := []struct {
		in       string
		expected string
	}{
		{"{\"foo\":\"bar\"}", "{\n  \"foo\": \"bar\"\n}\n"},
	}

	for _, c := range cases {
		in := strings.NewReader(c.in)
		buf := new(bytes.Buffer)
		DiLaurentis(in, buf, "  ")
		got := buf.String()
		if got != c.expected {
			t.Errorf("got %q but expected %q for input %q", got, c.expected, c.in)
		}
	}
}

func TestArgs(t *testing.T) {
	json := "{\"foo\":\"bar\"}"
	cases := []struct {
		argv     []string
		expected string
	}{
		{[]string{}, "{\n  \"foo\": \"bar\"\n}\n"},
		{[]string{"--indent", ""}, "{\n\"foo\": \"bar\"\n}\n"},
		{[]string{"--indent", "  "}, "{\n  \"foo\": \"bar\"\n}\n"},
		{[]string{"--indent", "    "}, "{\n    \"foo\": \"bar\"\n}\n"},
	}

	for _, c := range cases {
		in := strings.NewReader(json)
		buf := new(bytes.Buffer)

		run(c.argv, in, buf)

		got := buf.String()
		if got != c.expected {
			t.Errorf("got %q but expected %q with args %q", got, c.expected, c.argv)
		}
	}
}
