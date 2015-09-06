package main_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/f2prateek/DiLaurentis"
)

func TestDiLaurentis(t *testing.T) {
	cases := []struct {
		in       string
		expected string
	}{
		{"{\"foo\":\"bar\"}", "{\n  \"foo\": \"bar\"\n}"},
	}

	for _, c := range cases {
		in := strings.NewReader(c.in)
		buf := new(bytes.Buffer)
		main.DiLaurentis(in, buf)
		got := buf.String()
		if got != c.expected {
			t.Errorf("got %q but expected %q for input %q", got, c.expected, c.in)
		}
	}
}
