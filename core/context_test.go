package core_test

import (
	"context"
	"testing"
	_ "unsafe"

	. "github.com/adisubagja/xcore/core"
)

func TestFromContextPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expect panic, but nil")
		}
	}()

	MustFromContext(context.Background())
}
