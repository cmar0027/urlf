package urlf

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/google/uuid"
)

type test struct {
	format      string
	args        []any
	expected    string
	shouldPanic bool
}

func assertPanic(t *testing.T, i int, msg any, f func()) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("test %d failed: the code did not panic", i)
		} else if r != msg {
			t.Fatalf("test %d failed: the code panicked but with an unexpected message: '%v'", i, r)
		}
	}()
	f()
}

func TestSprintf(t *testing.T) {
	rawUUID := "00000000-0000-0000-0000-000000000000"
	id, err := uuid.Parse(rawUUID)
	if err != nil {
		panic("error parsing uuid")
	}

	tests := []test{
		{
			format:   "",
			args:     []any{},
			expected: "",
		},
		{
			format:   "a",
			args:     []any{},
			expected: "a",
		},
		{
			format:   "abc",
			args:     []any{},
			expected: "abc",
		},
		{
			format:   "a%%bc",
			args:     []any{},
			expected: "a%bc",
		},
		{
			format:   "abc%%",
			args:     []any{},
			expected: "abc%",
		},
		{
			format:      "abc%",
			args:        []any{},
			expected:    "found invalid '%' at index 3",
			shouldPanic: true,
		},
		{
			format:      "%p%",
			args:        []any{""},
			expected:    "found invalid '%' at index 2",
			shouldPanic: true,
		},
		{
			format:      "%q%",
			args:        []any{""},
			expected:    "found invalid '%' at index 2",
			shouldPanic: true,
		},
		{
			format:      "%%%",
			args:        []any{},
			expected:    "found invalid '%' at index 2",
			shouldPanic: true,
		},
		{
			format:   "/%p/",
			args:     []any{"hello-world"},
			expected: "/hello-world/",
		},
		{
			format:   "/%p/",
			args:     []any{id},
			expected: fmt.Sprintf("/%s/", url.PathEscape(id.String())),
		},
		{
			format:   "%p",
			args:     []any{"🫠"},
			expected: "%F0%9F%AB%A0",
		},
		{
			format:   "%q",
			args:     []any{"🫠"},
			expected: "%F0%9F%AB%A0",
		},
		{
			format:      "%q/%p",
			args:        []any{"hello", "world"},
			expected:    "found %p after %q",
			shouldPanic: true,
		},
		{
			format:   "%p/%q",
			args:     []any{"hello", "world"},
			expected: "hello/world",
		},
		{
			format:   "%aa",
			args:     []any{},
			expected: "%aa",
		},
		{
			format:      "%z",
			args:        []any{},
			expected:    "found invalid '%' at index 0",
			shouldPanic: true,
		},
		{
			format:      "%p",
			args:        []any{},
			expected:    "more arguments expected",
			shouldPanic: true,
		},
		{
			format:      "%p",
			args:        []any{"", ""},
			expected:    "given more arguments than expected",
			shouldPanic: true,
		},
	}

	for i, test := range tests {
		if test.shouldPanic {
			assertPanic(t, i, test.expected, func() {
				Sprintf(test.format, test.args...)
			})
		} else {
			v := Sprintf(test.format, test.args...)
			if v != test.expected {
				t.Fatalf("test %d failed: expected '%s' but got '%s'", i, test.expected, v)
			}
		}
	}

}
