package tools

import (
	"fmt"
	"os"
)

// DealMessage .
func DealMessage(ok bool, message string) {
	if !ok {
		fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}
}
