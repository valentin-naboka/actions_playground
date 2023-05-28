package foo_test

import (
	"actions_playground/foo"
	"testing"
)

func TestFoo(t *testing.T) {
	if foo.Bar() != "Hello world!" {
		t.Error("test failed")
	}
}
