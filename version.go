package main

import (
	"fmt"
	"runtime"
)

var (
	// These variables are initialized via the linker -X flag in the
	// top-level Makefile when compiling release binaries.
	Tag       = "unknown" // Tag of this build (git describe)
	Time      string      // Build time in UTC (year/month/day hour:min:sec)
	Revision  string      // SHA-1 of this build (git rev-parse)
	Platform  = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
	GoVersion = runtime.Version()
)
