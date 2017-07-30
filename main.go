/**
 * Copyright (c) 2017 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/edgexfoundry/exportclient/api"
	"github.com/edgexfoundry/exportclient/mongo"
	"github.com/cenkalti/backoff"
)

const (
	help string = `
Usage: exportclient [options]
Options:
	-a, --host	Host address
	-p, --port	Port
	-m, --mhost	Mongo host
	-q, --mport	Mongo port
	-h, --help	Prints this message end exits`

	httpHost = "0.0.0.0"
	httpPort = "7070"
	mongoHost = "0.0.0.0"
	mongoPort = "27017"
)

type (
	Opts struct {
		HTTPHost string
		HTTPPort string

		MongoHost string
		MongoPort string

		Help    bool
		Version bool
	}
)

var (
	opts Opts
)

func tryMongoInit() error {
    var err error

    log.Print("Connecting to MongoDB... ")
    err = mongo.InitMongo(opts.MongoHost, opts.MongoPort, opts.MongoHost)
    return err
}

func main() {
	flag.StringVar(&opts.HTTPHost, "a", httpHost, "HTTP server address.")
	flag.StringVar(&opts.HTTPHost, "host", httpHost, "HTTP server address.")
	flag.StringVar(&opts.HTTPPort, "p", httpPort, "HTTP server port.")
	flag.StringVar(&opts.HTTPPort, "port", httpPort, "HTTP server port.")
	flag.StringVar(&opts.MongoHost, "m", mongoHost, "MongoDB address.")
	flag.StringVar(&opts.MongoHost, "mhost", mongoHost, "MongoDB address.")
	flag.StringVar(&opts.MongoPort, "q", mongoPort, "MongoDB port.")
	flag.StringVar(&opts.MongoPort, "mport", mongoPort, "MongoDB port.")
	flag.BoolVar(&opts.Version, "version", false, "Print version information.")
	flag.BoolVar(&opts.Version, "v", false, "Print version information.")
	flag.BoolVar(&opts.Help, "h", false, "Show help.")
	flag.BoolVar(&opts.Help, "help", false, "Show help.")

	flag.Parse()

	if opts.Version {
		tw := tabwriter.NewWriter(os.Stdout, 2, 1, 2, ' ', 0)
		fmt.Fprintf(tw, "Build Tag:    %s\n", Tag)
		fmt.Fprintf(tw, "Build Time:   %s\n", Time)
		fmt.Fprintf(tw, "Platform:     %s\n", Platform)
		fmt.Fprintf(tw, "Go Version:   %s\n", GoVersion)
		fmt.Fprintf(tw, "Build SHA-1:  %s\n", Revision)
		tw.Flush()
		os.Exit(0)
	}

	if opts.Help {
		fmt.Printf("%s\n", help)
		os.Exit(0)
	}

	// Connect to MongoDB
	if err := backoff.Retry(tryMongoInit, backoff.NewExponentialBackOff()); err != nil {
		log.Fatalf("MongoDB: Can't connect: %v\n", err)
	} else {
		log.Println("MongoDB: Connected")
	}

	// Serve HTTP
	HTTPHost := fmt.Sprintf("%s:%s", opts.HTTPHost, opts.HTTPPort)
	http.ListenAndServe(HTTPHost, api.HTTPServer())
}
