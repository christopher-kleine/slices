package str_test

import (
	"strings"
	"testing"

	"github.com/christopher-kleine/slices/str"
)

func TestAll(t *testing.T) {
	src := []string{"foo", "bar", "hello", "world"}
	isLCase := func(v string) bool {
		return strings.ToLower(v) == v
	}
	has3Letters := func(v string) bool {
		return len(v) == 3
	}

	allLowerCase := str.All(src, isLCase)
	all3Letters := str.All(src, has3Letters)

	if !allLowerCase {
		t.Error("all entries in source should be lower-case")
	}

	if all3Letters {
		t.Error("some entries should have a different length than 3 letters")
	}
}
