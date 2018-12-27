package testutil

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func caller() string {
	const skip = 2 // caller()自身 + 各Assert関数をスキップ
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	file = file[strings.LastIndex(file, "/")+1:]
	// Backspace文字を使ってGoが出力する "assert.go:nn: " を打ち消すhack
	bs := "\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08"
	return fmt.Sprintf(bs+"%s:%d:", file, line)
}

// AssertEquals は実値と期待値が同値か判定を実施する
func AssertEquals(t *testing.T, title string, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("%s %s: unexpected, actual: `%v`, expected: `%v`", caller(), title, actual, expected)
	}
}

// AssertIntEquals は実値と期待値が同数値か判定を実施する
func AssertIntEquals(t *testing.T, title string, actual, expected int) {
	if actual != expected {
		// スライスの長さを比較するケースを考慮して、Fatalにしてある
		t.Fatalf("%s %s: unexpected, actual: `%d`, expected: `%d`", caller(), title, actual, expected)
	}
}
