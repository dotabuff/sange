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
	Convey("parses a replay", t, func() {
		parser := ParserFromFile("replays/test.dem")
		So(func() { parser.Parse() }, ShouldNotPanic)
	})
}
