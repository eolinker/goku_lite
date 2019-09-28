package utils

import "strings"

//TrimPrefixAll trimPerfixAll
func TrimPrefixAll(s string, r string) string {

	for strings.HasPrefix(s, r) {
		s = strings.TrimPrefix(s, r)
	}
	return s
}

//TrimSuffixAll trimSuffixAll
func TrimSuffixAll(s string, r string) string {

	for strings.HasSuffix(s, r) {
		s = strings.TrimSuffix(s, r)
	}
	return s

}
