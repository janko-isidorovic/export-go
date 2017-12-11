
.PHONY: build test vet prepare edgexclient edgexdistro edgexdistro_zmq docker

# Make exec targets phony to not track changes in go files. Compilation is fast
.PHONY: edgexclient edgexdistro edgexdistro_zmq

default: build

edgexclient:
	go build -o edgexclient cmd/client/main.go

edgexdistro:
	go build -o edgexdistro cmd/distro/main.go

edgexdistro_zmq:
	go build -o edgexdistro_zmq -tags zeromq cmd/distro/main.go

build: edgexclient edgexdistro edgexdistro_zmq

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
