package kendoui

import (
	"strings"
	"net/http"
	"net/url"
)

type RequestInput map[string]interface{}

func Input(r *http.Request) RequestInput {
	i := make(RequestInput)

	i.Parse(r.URL.Query())
	i.Parse(r.Form)

	return i
}

// filter[filters][0][field]=Title
// filter[filters][0][ignoreCase]=true
// filter[filters][0][operator]=startswith
// filter[filters][0][value]=the
// filter[logic]=and
func (i *RequestInput) Parse(vs url.Values) {
	for key, value := range vs {
		p := *i

		keys := strings.Split(key, "[")
		len := len(keys)
		key = keys[0]

		for i := 1; i < len; i++ {
			p1, ok := p[key]
			if !ok {
				p1 = make(RequestInput)
				p[key] = p1
			}
			p = (p1).(RequestInput)

			key = strings.TrimRight(keys[i], "]")
		}
		p[key] = value
	}
}
