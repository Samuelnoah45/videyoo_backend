package graphqlClient

import (
 "net/http"

)

type headersTransport struct {
 headers http.Header
 base    http.RoundTripper
}

func (t *headersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
 for k, v := range t.headers {
  req.Header.Set(k, v[0])
 }
 return t.base.RoundTrip(req) 
}

