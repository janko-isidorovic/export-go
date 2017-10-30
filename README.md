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

![edgy](https://lh6.googleusercontent.com/m4-1DvNQa1UwsOE4PK7TytR8is6eaVYEGVPwpeIl8EFjn2pgYGxQ0-AcjLFUJxF_UVHCORxPEiGCJs92pdJsxlK9M5T5_31-MHR3=w1920-h954)


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
