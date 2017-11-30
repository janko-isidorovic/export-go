//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

const (
	emptyRegistrationList = "[]"
	registrationStr       = `{"_id":"5a15918fa4a9b92af1c94bab","created":0,"modified":0,"origin":1471806386919,"name":"OTROMAS-1","addressable":{"Name":"OTROMAS-1","Method":"POST","Protocol":"TCP","Address":"127.0.0.1","Port":1883,"Path":"","Publisher":"FuseExportPublisher_OTROMAS-1","User":"dummy","Password":"dummy","Topic":"FuseDataTopic"},"format":"JSON","filter":{},"encryption":{},"compression":"NONE","enable":true,"destination":"MQTT_TOPIC"}`
	oneRegistrationList   = "[" + registrationStr + "]"
)

func TestClientRegistrationsEmpty(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, emptyRegistrationList)
	}

	// create test server with handler
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	regs := getRegistrationsURL(ts.URL)
	if regs == nil {
		t.Fatal("nil registration list")
	}
	if len(regs) != 0 {
		t.Fatal("Registration should be empty")
	}
}

func TestClientRegistrations(t *testing.T) {
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, oneRegistrationList)
	}

	// create test server with handler
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	regs := getRegistrationsURL(ts.URL)
	if regs == nil {
		t.Fatal("nil registration list")
	}
	if len(regs) != 1 {
		t.Fatal("Registration list should have only a registration")
	}
}
