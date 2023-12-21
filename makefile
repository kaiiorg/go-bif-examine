build:
	go build -o ./bin/ ./cmd/go-bif-examine

build-cli:
	go build -o ./bin/ ./cmd/go-bif-examine-cli

fmt:
	gofmt -w -s pkg cmd

grpc:
	./protoc/bin/protoc \
		--go_out=./pkg/rpc \
		--go-grpc_out=./pkg/rpc \
		./pkg/rpc/pb/bif_examine.proto

install_grpc_prereqs:
	mkdir -p ./protoc
	cd ./protoc && wget -nc  https://github.com/protocolbuffers/protobuf/releases/download/v25.1/protoc-25.1-linux-x86_64.zip
	cd ./protoc && unzip protoc-25.1-linux-x86_64.zip
	chmod +x ./protoc/bin/protoc
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	# cd ./protoc && wget -nc https://github.com/protocolbuffers/protobuf-javascript/releases/download/v3.21.2/protobuf-javascript-3.21.2-linux-x86_64.tar.gz
	# cd ./protoc && tar -xzvf protobuf-javascript-3.21.2-linux-x86_64.tar.gz
