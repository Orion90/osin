package osin

import (
	"encoding/json"
	"io"
	"net/http"
)

// OutputJSON encodes the Response to JSON and writes to the http.ResponseWriter
func OutputJSON(rs *Response, w io.Writer, r *http.Request) error {
	// Add headers

	if rs.Type == REDIRECT {
		q := w.(http.ResponseWriter)
		for i, k := range rs.Headers {
			for _, v := range k {
				q.Header().Add(i, v)
			}
		}
		// Output redirect with parameters
		u, err := rs.GetRedirectUrl()
		if err != nil {
			return err
		}
		q.Header().Add("Location", u)
		q.WriteHeader(302)
	} else {
		encoder := json.NewEncoder(w)
		err := encoder.Encode(rs.Output)
		if err != nil {
			return err
		}
	}
	return nil
}
