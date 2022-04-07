.PHONY: protos

protos:
	protoc --go_out=protos/currency --go_opt=paths=source_relative \
    --go-grpc_out=protos/currency --go-grpc_opt=paths=source_relative protos/currency.proto