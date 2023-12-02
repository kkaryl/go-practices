package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	// b := []byte(s)
	// for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
	// 	b[i], b[j] = b[j], b[i]
	// }
	// return string(b)

	// use runes instead of bytes: https://go.dev/blog/strings
	if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
	fmt.Printf("input: %q\n", s)
	r := []rune(s)
	fmt.Printf("runes: %q\n", r)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
