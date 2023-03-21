# generate the messages
# generate the services (--go-grpc_out=./)
gen-cal:
	protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false calculatorpb/calculator.proto
run-server:
	go run server/server.go
run-client:
	go run client/client.go

