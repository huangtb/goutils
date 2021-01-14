package goutils

import (
	"fmt"
	"testing"
)

func TestFormatTimestamp(t *testing.T) {
	fmt.Println(FormatTimestamp(1610594931111,YYYY_mm_DDHH_mm_SS))
}
