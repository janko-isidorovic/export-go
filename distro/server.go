//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
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

func replyNotifyRegistrations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "")
	RefreshRegistrations()
}

// HTTPServer function
func HTTPServer() http.Handler {
	mux := bone.New()

	mux.Get("/api/v1/ping", http.HandlerFunc(replyPing))
	mux.Put("/api/v1/notify/registrations", http.HandlerFunc(replyNotifyRegistrations))

	return mux
}
