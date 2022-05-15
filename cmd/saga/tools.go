//go:build tools
// +build tools

package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "go.einride.tech/protoc-gen-go-cli"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
