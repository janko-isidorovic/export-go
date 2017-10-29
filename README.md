# EdgeX Foundry Export Services
[![Go Report Card](https://goreportcard.com/badge/github.com/drasko/edgex-export)](https://goreportcard.com/report/github.com/drasko/edgex-export)
[![license](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)

Go implementation of EdgeX Export services.

[export-client](https://github.com/edgexfoundry/export-client),
[export-distro](https://github.com/edgexfoundry/export-distro) and
[export-domain](https://github.com/edgexfoundry/export-domain) have been merged
into one single repo. Additionaly, some classes from
[core-domain](https://github.com/edgexfoundry/core-domain) have been duplicated
into this repo (needed for linking and data isolation of microservices).

Repo contains two microservices, `export-client` and `export-distro`.

![edgy](https://lh6.googleusercontent.com/xTd8zK0zT6-b9J_TL5vMm6tPGWlPeVYs8zRx4S0gt-qy4bmaE7ULyOspZ3lKkwsjIV8JVN1S93-0DuKa8wZj0F30V7pmm4kZyVzK=w1920-h903)

## Install and Deploy

Currently only `export-client` is functional.

To fetch the code and start the microservice execute:

```
go get github.com/drasko/edgex-export
cd $GOPATH/src/github.com/drasko/edgex-export
glide install
go run cmd/client/main.go
```
## Community
- Chat: https://chat.edgexfoundry.org/home
- Mainling lists: https://lists.edgexfoundry.org/mailman/listinfo

## License
[Apache-2.0](LICENSE)
