/*
Copyright 2019

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Written by Muhammet Arslan <github.com/geass>, December 2019
*/

package http

import (
	"time"

	"github.com/valyala/fasthttp"
)

// Option is the func interface to assign options
type Option func(*Options)

// Options is a struct for HTTP Server Options
type Options struct {
	Address string

	Handler fasthttp.RequestHandler

	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	MaxConnsPerIP        int
	MaxRequestsPerConn   int
	MaxKeepaliveDuration time.Duration
}

// WithAddress defines the HTTPServer Listening Address
func WithAddress(address string) Option {
	return func(o *Options) {
		o.Address = address
	}
}

// WithHandler defines the HTTPServer Handler
func WithHandler(handler fasthttp.RequestHandler) Option {
	return func(o *Options) {
		o.Handler = handler
	}
}

// WithReadTimeout defines the HTTPServer read timeout
func WithReadTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = timeout
	}
}

// WithWriteTimeout defines the HTTPServer write timeout
func WithWriteTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = timeout
	}
}

// WithMaxConnsPerIP defines the HTTPServer maximum connections limit per IP
func WithMaxConnsPerIP(limit int) Option {
	return func(o *Options) {
		o.MaxConnsPerIP = limit
	}
}

// WithMaxRequestsPerConn defines the HTTPServer maximum requrest limit per connection
func WithMaxRequestsPerConn(limit int) Option {
	return func(o *Options) {
		o.MaxRequestsPerConn = limit
	}
}

// WithMaxKeepaliveDuration defines the HTTPServer maximum keep-alive duration
func WithMaxKeepaliveDuration(timeout time.Duration) Option {
	return func(o *Options) {
		o.MaxKeepaliveDuration = timeout
	}
}
