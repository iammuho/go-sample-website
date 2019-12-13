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

package main

import (
	"github.com/geass/go-sample-website/templates/index"
	"github.com/geass/go-sample-website/templates/layout"

	"github.com/valyala/fasthttp"
)

// Index function renders the dashboard index page
func Index() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.SetContentType("text/html")

		p := &index.IndexPage{
			CTX: ctx,
		}
		layout.WriteBaseLayout(ctx, p)
	}
}
