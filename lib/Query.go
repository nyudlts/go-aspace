package lib

import "fmt"

type QueryString struct {
	Query string
}

func (q *QueryString)AddParameter(code string, value string) {
	if(q.Query == "") {
		q.Query = fmt.Sprintf("%s=%s", code, value)
	} else {
		q.Query = fmt.Sprintf("%s&%s=%s", q.Query, code, value)
	}
}
