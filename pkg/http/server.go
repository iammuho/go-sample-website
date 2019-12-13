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
	"github.com/valyala/fasthttp/reuseport"
)

var defaultOptions = Options{
	Address: "0.0.0.0:6060",

	ReadTimeout:          5 * time.Second,
	WriteTimeout:         10 * time.Second,
	MaxConnsPerIP:        500,
	MaxRequestsPerConn:   500,
	MaxKeepaliveDuration: 5 * time.Second,
}

// Server is a structure which exposes FastHTTP library
type Server struct {
	Options *Options
	Server  *fasthttp.Server
	Handler fasthttp.RequestHandler
}

// New instantiate new logging function by using ubers zap library
// for the development environments
func New(opts ...Option) *Server {

	options := defaultOptions
	for _, o := range opts {
		o(&options)
	}

	// Return the SERVER
	return &Server{
		Options: &options,
		Server: &fasthttp.Server{
			Handler:              options.Handler,
			ReadTimeout:          options.ReadTimeout,
			WriteTimeout:         options.WriteTimeout,
			MaxConnsPerIP:        options.MaxConnsPerIP,
			MaxRequestsPerConn:   options.MaxRequestsPerConn,
			MaxKeepaliveDuration: options.MaxKeepaliveDuration,
		},
	}
}

// Start starts fasthttp application server
func (s *Server) Start() error {

	ln, err := reuseport.Listen("tcp4", s.Options.Address)
	if err != nil {
		return err
	}

	return s.Server.Serve(ln)
}

// Stop stops http application server
func (s *Server) Stop() error {
	return s.Server.Shutdown()
}
