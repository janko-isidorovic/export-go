package api

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
	//mux.Get("/api/v1/registration/reference/:type", http.HandlerFunc(getRefListHandler))
	mux.Get("/api/v1/registration", http.HandlerFunc(getAllReg))
	//mux.Get("/api/v1/registration/name/:name", http.HandlerFunc(getRegByNameHandler))
	//mux.Post("/api/v1/registration", http.HandlerFunc(addRegHandler))
	//mux.Put("/api/v1/registration", http.HandlerFunc(updateRegHandler))
	//mux.Delete("/api/v1/registration/id/:id", http.HandlerFunc(delRegByIDHandler))
	//mux.Delete("/api/v1/registration/name/:name", http.HandlerFunc(delRegByNameHandler))

	return mux
}
