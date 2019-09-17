package utils

import "strings"

func TrimPrefixAll(s string, r string) string {

	for strings.HasPrefix(s, r) {
		s = strings.TrimPrefix(s, r)
	}
	return s
}
func TrimSuffixAll(s string, r string) string {

	for strings.HasSuffix(s, r) {
		s = strings.TrimSuffix(s, r)
	}
	return s

}
