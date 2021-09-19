TARGET_DIR = helloworld

build: $(TARGET_DIR)
	protoc --lint_out=. $(TARGET_DIR)/*.proto
	protoc -I . --go_out . --go_opt=paths=source_relative --go-grpc_out . --go-grpc_opt=paths=source_relative $(TARGET_DIR)/*.proto
