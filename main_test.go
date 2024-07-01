package main_test

import (
	"fmt"
	"testing"
)

func testMain(t *testing.T) {
	testStrings := []string{
		"etum.wtf",
		"etum.wtf/",
		"",
	}
	for _, str := range testStrings {
		if want != "/" {
			fmt.Println(str)
		}
	}
}
