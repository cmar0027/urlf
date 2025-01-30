// Package urlf provides fmt-like formatting for URLs
package urlf

import (
	"fmt"
	"net/url"
	"strings"
)

type arg struct {
	identifier byte
	index      int
}

func findArgs(format string) (args []arg) {
	args = []arg{}

	for i := 0; i < len(format)-1; {
		if format[i] == '%' {
			if format[i+1] == 'p' {

			} else if format[i+1] == 'q' {

			} else if format[i+1] == '%' {

			} else {

			}
		} else {

		}
	}
	return
}

func isEsadecimal(v byte) bool {
	return (48 <= v && v <= 57) || (65 <= v && v <= 70) || (97 <= v && v <= 102)
}

// Sprintf securely formats a url.
// Works like fmt.Sprintf but only accepts %p, %q and %%:
// - %p calls url.PathEscape on the corresponding argument
// - %q calls url.QueryEscape on the corresponding argument
// - %% escapes an '%'
func Sprintf(format string, a ...any) string {

	builder := strings.Builder{}

	lastStop := 0
	argIndex := 0

	firstQueryFound := false
	i := 0
	for i < len(format)-1 {
		if format[i] == '%' {
			switch format[i+1] {
			case 'p':
				if firstQueryFound {
					panic("found %p after %q")
				}
				builder.WriteString(format[lastStop:i])
				if argIndex >= len(a) {
					panic("more arguments expected")
				}
				builder.WriteString(url.PathEscape(fmt.Sprint(a[argIndex])))
				argIndex += 1
				lastStop = i + 2
				i += 2
			case 'q':
				firstQueryFound = true
				builder.WriteString(format[lastStop:i])
				if argIndex >= len(a) {
					panic("more arguments expected")
				}
				builder.WriteString(url.QueryEscape(fmt.Sprint(a[argIndex])))
				argIndex += 1
				lastStop = i + 2
				i += 2
			case '%':
				builder.WriteString(format[lastStop:i])
				builder.WriteByte('%')
				lastStop = i + 2
				i += 2
			default:
				if i < len(format)-2 && isEsadecimal(format[i+1]) && isEsadecimal(format[i+2]) {
					i += 3
				} else {
					panic(fmt.Sprintf("found invalid '%%' at index %d", i))
				}
			}
		} else {
			i += 1
		}
	}

	if argIndex != len(a) {
		panic("given more arguments than expected")
	}

	if len(format) > 0 && i < len(format) && format[len(format)-1] == '%' {
		panic(fmt.Sprintf("found invalid '%%' at index %d", len(format)-1))
	}

	if lastStop < len(format) {
		builder.WriteString(format[lastStop:])
	}

	return builder.String()
}
