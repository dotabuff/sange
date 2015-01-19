package sange

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	spew.Config.SortKeys = true
}

func TestParser(t *testing.T) {
	Convey("parses replay", t, func() {
		parser := ParserFromFile("replays/test.dem")
		So(func() { parser.Parse() }, ShouldNotPanic)
	})
	Convey("parses replay1", t, func() {
		parser := ParserFromFile("replays/test1.dem")
		So(func() { parser.Parse() }, ShouldNotPanic)
	})
	Convey("parses replay2", t, func() {
		parser := ParserFromFile("replays/test2.dem")
		So(func() { parser.Parse() }, ShouldNotPanic)
	})
}
