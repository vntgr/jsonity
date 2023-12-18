package main

import (
	"fmt"
)

// Options is Pretty options
type Options struct {
	// Width is an max column width for single line arrays
	// Default is 80
	Width int
	// Prefix is a prefix for all lines
	// Default is an empty string
	Prefix string
	// Indent is the nested indentation
	// Default is two spaces
	Indent string
	// SortKeys will sort the keys alphabetically
	// Default is false
	SortKeys bool
}

// DefaultOptions is the default options for pretty formats.
var DefaultOptions = &Options{Width: 80, Prefix: "", Indent: "  ", SortKeys: false}

func Pretty(json []byte) []byte { return PrettyOptions(json, nil) }

func PrettyOptions(json []byte, opts *Options) []byte {
	if opts == nil {
		opts = DefaultOptions
	}
	buf := make([]byte, 0, len(json))

	if len(opts.Prefix) != 0 {
		buf = append(buf, opts.Prefix...)
	}

	if len(buf) > 0 {
		buf = append(buf, '\n')
	}
	return buf
}

var jsn = `{"name":  {"first":"Tom","last":"Anderson"},  "age":37,
"children": ["Sara","Alex","Jack"],
"fav.movie": "Deer Hunter", "friends": [
    {"first": "Janet", "last": "Murphy", "age": 44}
  ]}`

func main() {
	fmt.Println(jsn)
	opts := &Options{
		Prefix: "nnn",
	}
	result := PrettyOptions([]byte(jsn), opts)
	fmt.Println(string(result))
}
