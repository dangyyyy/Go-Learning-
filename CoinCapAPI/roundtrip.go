package Coincap

import (
	"fmt"
	"io"
	"net/http"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "%s %s %s\n", req.Method, req.URL, req.Proto)
	return l.next.RoundTrip(req)
}
