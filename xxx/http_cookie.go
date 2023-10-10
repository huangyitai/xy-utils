package xxx

import "net/http"

// CookiesToJSONStr 将cookie切片学序列化为json的方法
func CookiesToJSONStr(cookies []*http.Cookie) string {
	m := map[string]string{}
	for _, cookie := range cookies {
		if cookie != nil {
			m[cookie.Name] = cookie.String()
		}
	}
	return ToJSONStr(m)
}
