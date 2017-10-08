# EdgeX Foundry Export Services

Go implementation of EdgeX Export services.

[export-client](https://github.com/edgexfoundry/export-client),
[export-distro](https://github.com/edgexfoundry/export-distro) and
[export-domain](https://github.com/edgexfoundry/export-domain) have been merged
into one single repo.

Repo contains two microservices, `export-client` and `export-distro`

# Install and Deploy

Currently only `export-client` is functional.

To fetch the code and start the microservice execute:

```
go get github.com/drasko/edgex-export
cd $GOPATH/src/github.com/drasko/edgex-export/cmd/client
go run main.go
```

# License
[Apache-2.0](LICENSE)
