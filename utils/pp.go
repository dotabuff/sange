package utils

import (
	"runtime"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var (
	P = spew.Dump
)

// use this to make debugging output that shows the location.
func PP(args ...interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		f := runtime.FuncForPC(pc)
		fParts := strings.Split(f.Name(), ".")
		fun := fParts[len(fParts)-1]
		s := spew.Sprintf("vvvvvvvvvvvvvvv %s vvvvvvvvvvvvvvv\n", fun)
		spew.Print(s)
		spew.Dump(args...)
		spew.Println(strings.Repeat("^", len(s)-1))
	} else {
		spew.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
		spew.Dump(args...)
		spew.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	}
}
