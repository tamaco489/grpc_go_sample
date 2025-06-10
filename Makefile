# ========================================================================
# setup
# ========================================================================
#
# Usage:
#   make install-tools
#   make proto-gen

.PHONY: install-tools proto-gen

# Install dependencies
install-tools:
	@echo "Installing dependencies..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	@echo "Done!"

# Generate protobuf files
proto-gen:
	@echo "Generating protobuf files..."
	protoc \
		--go_out=./api/internal/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=./api/internal/gen \
		--go-grpc_opt=paths=source_relative \
		proto/healthcheck/healthcheck.proto
	@echo "Done!"


# ========================================================================
# grpc server request
# ========================================================================
#
# Usage:
#   make healthcheck
#
healthcheck:
	grpcurl -plaintext localhost:50051 healthcheck.HealthCheckService/GetHeartbeat
