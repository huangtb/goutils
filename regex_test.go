package goutils

import (
	"fmt"
	"testing"
)

func TestIsMatchRegex(t *testing.T) {
	fmt.Println(IsMatchRegex("a beta b","\\sbet\\s"))
}
