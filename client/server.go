//
// Copyright 2017 Mainflux.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package client

import (
	"net/http"

	"github.com/go-zoo/bone"
)

// HTTPServer function
func HTTPServer() http.Handler {
	mux := bone.New()

	// Status
	mux.Get("/status", http.HandlerFunc(getStatus))

	// Registration
	mux.Get("/api/v1/registration/:id", http.HandlerFunc(getRegByID))
	mux.Get("/api/v1/registration/reference/:type", http.HandlerFunc(getRegList))
	mux.Get("/api/v1/registration", http.HandlerFunc(getAllReg))
	mux.Get("/api/v1/registration/name/:name", http.HandlerFunc(getRegByName))
	mux.Post("/api/v1/registration", http.HandlerFunc(addReg))
	mux.Put("/api/v1/registration", http.HandlerFunc(updateReg))
	mux.Delete("/api/v1/registration/id/:id", http.HandlerFunc(delRegByID))
	mux.Delete("/api/v1/registration/name/:name", http.HandlerFunc(delRegByName))

	return mux
}
