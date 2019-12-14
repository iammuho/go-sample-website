// Copyright 2019 Muhammet Arslan <github.com/geass>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

//go:generate qtc -dir=templates -ext=html

import (
	"sync"

	"github.com/geass/go-sample-website/config"
	"github.com/geass/go-sample-website/pkg/http"
	"github.com/geass/go-sample-website/pkg/logger"

	"go.uber.org/zap"
)

func main() {

	// Initialize logger
	logger := logger.New(config.Config.Application.Name, config.Config.Application.Environment)

	// Debug Log
	logger.Debug("Debug! This will not be seen if environment is 'production'")

	// Register the router
	handler := registerRouter()

	// Implement the HTTP Server
	httpServer := http.New(
		http.WithAddress(config.Config.HTTPServer.Listen),
		http.WithHandler(handler),
		http.WithReadTimeout(config.Config.HTTPServer.ReadTimeout),
		http.WithWriteTimeout(config.Config.HTTPServer.WriteTimeout),
		http.WithMaxConnsPerIP(config.Config.HTTPServer.MaxConnsPerIP),
		http.WithMaxRequestsPerConn(config.Config.HTTPServer.MaxRequestsPerConn),
		http.WithMaxKeepaliveDuration(config.Config.HTTPServer.MaxKeepaliveDuration),
	)

	// Start the HTTP server
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		logger.Info(
			"HTTP Server is running...",
			zap.String("Address", config.Config.HTTPServer.Listen),
		)

		defer wg.Done()
		httpServer.Start()
	}()

	wg.Wait()
}
