.PHONY: protobufs

prefix := github.com/manujelko/go-grpc-microservices/pkg

protobufs:
	protoc --proto_path=protobufs \
	--go_out=pkg --go_opt=module=$(prefix) \
	--go-grpc_out=pkg --go-grpc_opt=module=$(prefix) \
	protobufs/currency.proto
