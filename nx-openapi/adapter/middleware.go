package adapter

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type (
	Middleware interface {
		WrapClient(*http.Client)
	}
	MiddlewareFunc func(*http.Request, func(*http.Request) (*http.Response, error)) (*http.Response, error)

	middleware struct {
		base http.RoundTripper
		f    MiddlewareFunc
	}
)

const (
	headerApiKey = "x-nxopen-api-key"
	ocidParam    = "ocid"
	gidParam     = "oguild_id"
)

func AddMiddlewaresToClient(c *http.Client, middlewares ...Middleware) {
	if c == nil {
		panic("attempted to add middleware to nil client")
	}

	if c.Transport == nil {
		c.Transport = http.DefaultTransport
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		m := middlewares[i]
		if m == nil {
			continue
		}
		m.WrapClient(c)
	}
}

func (m *middleware) WrapClient(c *http.Client) {
	m.base = c.Transport
	c.Transport = m
}

func (m *middleware) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.f(req, m.base.RoundTrip)
}

func APIKeyHeaderMiddleware(apiKey string) Middleware {
	return &middleware{
		f: func(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
			req.Header.Add(headerApiKey, apiKey)
			return next(req)
		},
	}
}

func OCIDMiddleware(ocid *string) Middleware {
	return &middleware{
		f: func(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
			if ocid == nil || len(*ocid) == 0 {
				return next(req)
			}

			u := req.URL
			q := u.Query()
			q.Add(ocidParam, *ocid)
			u.RawQuery = q.Encode()
			return next(req)
		},
	}
}

func GIDMiddleware(gid *string) Middleware {
	return &middleware{
		f: func(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
			if gid == nil || len(*gid) == 0 {
				return next(req)
			}

			u := req.URL
			q := u.Query()
			q.Add(gidParam, *gid)
			u.RawQuery = q.Encode()
			return next(req)
		},
	}
}

func ThrottleMiddleware(maxRps int) Middleware {
	if maxRps == 0 {
		return nil
	}
	return &middleware{
		f: func(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
			time.Sleep(time.Second / time.Duration(maxRps))
			return next(req)
		},
	}
}

func RetryMiddleware(retryCount int) Middleware {
	return &middleware{
		f: func(req *http.Request, next func(*http.Request) (*http.Response, error)) (res *http.Response, err error) {
			counter := retryCount
			for counter >= 0 {
				res, err = next(req)
				if err != nil || (res.StatusCode < http.StatusInternalServerError && res.StatusCode != http.StatusTooManyRequests) {
					break
				}
				counter -= 1
			}
			return
		},
	}
}

func Convert400ResponseToError() Middleware {
	return &middleware{
		f: func(req *http.Request, next func(*http.Request) (*http.Response, error)) (res *http.Response, err error) {
			res, err = next(req)
			if err != nil || res == nil || res.StatusCode < http.StatusBadRequest {
				return
			}
			b, _ := io.ReadAll(res.Body)
			k := string(b)
			return nil, errors.New(k)
		},
	}
}
