package utils

import "strings"

//TrimPrefixAll 删除所有前缀
func TrimPrefixAll(s string, r string) string {

	for strings.HasPrefix(s, r) {
		s = strings.TrimPrefix(s, r)
	}
	return s
}

//TrimSuffixAll 删除所有后缀
func TrimSuffixAll(s string, r string) string {

	for strings.HasSuffix(s, r) {
		s = strings.TrimSuffix(s, r)
	}
	return s

}
