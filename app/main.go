package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}
}

func matchLine(line []byte, pattern string) (bool, error) {
	if pattern == "\\d"{
		for _,value := range line {
			if isDigit(value){
				return true,nil
			}
		}
		return false,nil
	}else{
		
		if utf8.RuneCountInString(pattern) != 1 {
			return false, fmt.Errorf("unsupported pattern: %q", pattern)
		}

		ok := bytes.ContainsAny(line, pattern)
		return ok, nil
	}

}

func isDigit(value byte) bool {
	numbers := []byte{0x30,0x31,0x32,0x33,0x34,0x35,0x36,0x37,0x38,0x39}
	return bytes.ContainsAny(numbers,string(value))
}
