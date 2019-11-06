package interpreter

import "net/http"

//MergeBodys mergeBodys
func MergeBodys(bodys []interface{}) interface{} {

	if isAllMap(bodys) {

		b1 := bodys[0]
		mall := b1.(map[string]interface{})

		for _, b := range bodys {

			m := b.(map[string]interface{})

			for k, v := range m {
				mall[k] = v
			}

		}
		return mall
	}
	if isAllSlice(bodys) {

		return bodys
	}
	return make(map[string]interface{}, 0)
}

func isAllSlice(bodys []interface{}) bool {

	for _, b := range bodys {

		switch b.(type) {
		case []map[string]interface{}, []interface{}:
			continue
		default:
			return false
		}
	}
	return true
}
func isAllMap(bodys []interface{}) bool {

	for _, b := range bodys {

		switch b.(type) {
		case map[string]interface{}:
			continue
		default:
			return false
		}
	}
	return true
}

//MergeHeaders mergeHeaders
func MergeHeaders(Headers []http.Header) http.Header {

	header := Headers[0]

	for _, h := range Headers[1:] {

		for k, v := range h {

			header[k] = v
		}
	}

	return header
}

//MergeCookies mergeCookies
func MergeCookies(cookies []_Cookies) []*http.Cookie {

	allCookies := make(map[string]*http.Cookie)

	for _, cs := range cookies {
		for _, c := range cs {
			allCookies[c.Name] = c
		}
	}
	newCookies := make([]*http.Cookie, 0, len(allCookies))
	for _, c := range allCookies {

		newCookies = append(newCookies, c)

	}
	return newCookies
}
