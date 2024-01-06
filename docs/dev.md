# Development Setup
Assumes you're using some flavor of Ubuntu Linux or WSL

1. Install Go 1.21.*
2. Install make
3. Install gprc prereqs: `make install_grpc_prereqs`

## Build Binaries
1. `make` (or `make build`)
2. `make build-cli`
3. `make build-whisperer`

## Update gRPC Generated Code
1. `make grpc`
