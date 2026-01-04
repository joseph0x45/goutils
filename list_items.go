package goutils

import (
	"fmt"
	"strings"
)

func pad(s string, width int) string {
	return s + strings.Repeat(" ", width-len(s))
}

func ListItems[T any](
	items []T,
	headers []string,
	extractors []func(T) string,
	padding int,
) {
	if len(headers) != len(extractors) {
		panic("headers and extractors length mismatch")
	}

	cols := len(headers)
	widths := make([]int, cols)

	for i, h := range headers {
		widths[i] = len(h)
	}

	for _, item := range items {
		for i, f := range extractors {
			v := f(item)
			if len(v) > widths[i] {
				widths[i] = len(v)
			}
		}
	}

	for i := range widths {
		widths[i] += padding
	}

	for i, h := range headers {
		fmt.Print(pad(h, widths[i]))
	}
	fmt.Println()

	for _, w := range widths {
		fmt.Print(strings.Repeat("-", w))
	}
	fmt.Println()

	for _, item := range items {
		for i, f := range extractors {
			fmt.Print(pad(f(item), widths[i]))
		}
		fmt.Println()
	}
}
