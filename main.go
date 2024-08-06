package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
)

var (
	decode = flag.Bool("decode", true, "Decode url encoded value")
	encode = flag.Bool("encode", false, "Encode as url encoded value")
)

func main() {
	flag.Parse()

	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read from stdin: %v\n", err)
		os.Exit(1)
	}

	if *encode {
		fmt.Fprintf(os.Stdout, "%s\n", url.QueryEscape(string(raw)))
		os.Exit(0)
	} else {
		dec, err := url.QueryUnescape(string(raw))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to unescape input: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "%s\n", dec)
		os.Exit(0)
	}
}
