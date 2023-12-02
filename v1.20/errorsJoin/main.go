package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var ByeError = errors.New("BYE!")

func sillyMethod(s string) error {
	if !utf8.ValidString(s) {
        return errors.New("input is not valid UTF-8")
    } else if s == "HELLO!" {
		return ByeError
	}
	return nil
}

func main() {
	validString := []string{"Hello, world!", "HELLO!", "Ŋ", "\x8a", "\x8b", "HELLO!"}
	var err error
	for _, s := range validString {
		err = errors.Join(err, sillyMethod(s))
	}
	fmt.Println(err)
	if errors.Is(err, ByeError) {
		fmt.Println("error check!")
	}
}