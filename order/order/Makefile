#######################  VARS  ##########################
baseDir = $$(basename $$PWD)
baseProto = internal/proto/$(baseDir)/$(baseDir).proto

#######################  GO  ##########################
run:
	go run main.go &

ci-cd: test-clean test \
	run

#######################  PROTO  ##########################
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(baseProto)


#######################  TEST  ##########################
check: test-clean test

test-clean:
	go clean -testcache

test:
	go test -v ./tests/...
