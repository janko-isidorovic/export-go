//
// Copyright (c) Mainflux
//
// Mainflux server is licensed under an Apache license, version 2.0.
// All rights not explicitly granted in the Apache license, version 2.0 are reserved.
// See the included LICENSE file for more details.
//

package distro

import (
	"io"
	"net/http"

	"github.com/go-zoo/bone"
)

func replyPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	str := `pong`
	io.WriteString(w, str)
}

// HTTPServer function
func HTTPServer() http.Handler {
	mux := bone.New()

	mux.Get("/api/v1/ping", http.HandlerFunc(replyPing))

	return mux
}
