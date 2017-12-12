# Copyright 2017 Cavium
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#


.PHONY: build test vet prepare edgexclient edgexdistro edgexdistro_zmq docker

# Make exec targets phony to not track changes in go files. Compilation is fast
.PHONY: client distro distro_zmq

default: build

client:
	go build -o client cmd/client/main.go

distro:
	go build -o distro cmd/distro/main.go

distro_zmq:
	go build -o distro_zmq -tags zeromq cmd/distro/main.go

build: client distro distro_zmq

docker:
	docker build -f Dockerfile.client  .
	docker build -f Dockerfile.distro  .

test:
	go test `glide novendor`

vet:
	go vet `glide novendor` 

coverage:
	go test -covermode=count -coverprofile=cov.out ./distro
	go tool cover -html=cov.out -o distroCoverage.html
	go test -covermode=count -coverprofile=cov.out ./client
	go tool cover -html=cov.out -o clientCoverage.html
	rm cov.out

bench:
	go test -run=XXX -bench=. ./distro

profile:
	go test -run=XXX -bench=.  -cpuprofile distro.cpu ./distro
	go test -run=XXX -bench=.  -memprofile distro.mem ./distro

prepare:
	glide install

clean:
	rm -f client distro distro_zmq cov.out distroCoverage.html \
       clientCoverage.html distro.cpu distro.mem
