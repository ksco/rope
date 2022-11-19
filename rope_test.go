package rope

import (
	"runtime/debug"
	"testing"
)

func assert(t *testing.T, condition bool) {
	if !condition {
		debug.PrintStack()
		t.Fail()
	}
}

func TestNewRope(t *testing.T) {
	r := NewRopeFromStr("Hell")
	r = Append(r, NewRope([]byte(",")))
	r = AppendBytes(r, []byte(" "))
	r = AppendStr(r, "World")
	r = Insert(r, 4, NewRopeFromStr("o"))

	full := "Hello, World"
	assert(t, Str(r) == full)
	assert(t, Len(r) == len(full))
	assert(t, Str(Sub(r, 3, 10)) == full[3:10])
}
