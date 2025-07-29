package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(target string) http.Handler {
	remote, err := url.Parse(target)
	if err != nil {
		panic("invalid proxy target: " + target)
	}
	return httputil.NewSingleHostReverseProxy(remote)
}
