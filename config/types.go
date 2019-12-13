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

package config

import "time"

// Config stores configuration values
var Config *config

type config struct {

	// Application provides the application configurations.
	Application struct {
		Name        string `env:"GOSAMPLE_APPLICATION_NAME"     envDefault:"go-sample-website"`
		Environment string `env:"GOSAMPLE_APPLICATION_ENVIRONMENT"     envDefault:"development"`
		Version     string `env:"GOSAMPLE_APPLICATION_VERSION"     envDefault:"1.0"`
	}

	// HTTPServer provides the HTTP server configuration.
	HTTPServer struct {
		Listen string `env:"GOSAMPLE_SERVER_LISTEN"     envDefault:"0.0.0.0:6060"`

		ReadTimeout          time.Duration `env:"GOSAMPLE_SERVER_READ_TIMEOUT" envDefault:"5s"`
		WriteTimeout         time.Duration `env:"GOSAMPLE_SERVER_WRITE_TIMEOUT" envDefault:"5s"`
		MaxConnsPerIP        int           `env:"GOSAMPLE_SERVER_MAX_CONN_PER_IP" envDefault:"50"`
		MaxRequestsPerConn   int           `env:"GOSAMPLE_SERVER_MAX_REQUESTS_PER_CONN" envDefault:"10"`
		MaxKeepaliveDuration time.Duration `env:"GOSAMPLE_SERVER_MAX_KEEP_ALIVE_DURATION" envDefault:"5s"`
	}
}
