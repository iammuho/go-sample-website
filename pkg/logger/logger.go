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

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a log structure which exposes Uber/ZAP library
type Logger struct{ *zap.Logger }

// New instantiate new logging function by using ubers zap library
// for the development environments
func New(name string, environment string) *Logger {

	logger, _ := newDevelopmentLogger()

	if environment == "production" {
		logger, _ = zap.NewProduction()
	}

	logger = logger.Named(name)

	defer logger.Sync()

	return &Logger{logger}
}

// newDevelopmentLogger will setup a new Development Logger
func newDevelopmentLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return cfg.Build(zap.AddCaller())
}
