protoc-gen-grpc-authority:
	protoc -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative services/authority/proto/authority.proto

protoc-gen-grpc-customer:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative services/customer/proto/customer.proto

protoc-gen-grpc-gateway:
	protoc -I. -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	--go_out ./ --go_opt paths=source_relative \
	--go-grpc_out ./ --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./ --grpc-gateway_opt paths=source_relative \
	./services/gateway/proto/gateway.proto



