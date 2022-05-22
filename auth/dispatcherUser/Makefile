#######################  VARS  ##########################
baseDir = $$(basename $$PWD)
baseProto = proto/$(baseDir)/$(baseDir).proto

proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ${arg}

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(baseProto)

ci-cd: test-clean test \
	run

check: test-clean test

test-clean:
	go clean -testcache

test:
	go test -v ./tests/...

run:
	go run main.go &