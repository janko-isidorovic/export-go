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
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetStatus(t *testing.T) {
	cases := []struct {
		body string
		code int
	}{
		{`{"running": true}`, 200},
	}

	url := ts.URL + "/status"

	for i, c := range cases {
		res, err := http.Get(url)
		if err != nil {
			t.Errorf("case %d: %s", i+1, err.Error())
		}

		if res.StatusCode != c.code {
			t.Errorf("case %d: expected status %d got %d", i+1, c.code, res.StatusCode)
		}

		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Fatalf("case %d: %s", i+1, err.Error())
		}

		if c.body != string(body) {
			t.Errorf("case %d: expected response %s got %s", i+1, c.body, string(body))
		}
	}
}
