gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/permissions.proto

run:
	go run ./cmd/permissionsservice