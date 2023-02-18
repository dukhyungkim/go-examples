TARGET_DIR = order

build: $(TARGET_DIR)
	protoc --lint_out=. $(TARGET_DIR)/*.proto
	protoc -I . --go_out . --go_opt=paths=source_relative --go-grpc_out . --go-grpc_opt=paths=source_relative  --grpc-gateway_out . --grpc-gateway_opt paths=source_relative $(TARGET_DIR)/*.proto

gen: $(TARGET_DIR)
	protoc --lint_out=. $(TARGET_DIR)/*.proto
	protoc -I . --openapiv2_out ./swagger --openapiv2_opt logtostderr=true $(TARGET_DIR)/*.proto

clean: $(TARGET_DIR)
	rm -rf $(TARGET_DIR)/*.go
