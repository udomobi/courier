package utils

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func ToAscii(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	transformer := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	str, _, _ = transform.String(transformer, str)
	return str
}