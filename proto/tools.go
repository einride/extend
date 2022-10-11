//go:build tools
// +build tools

package tools

import (
	_ "go.einride.tech/aip/cmd/protoc-gen-go-aip"
	_ "go.einride.tech/iam/cmd/protoc-gen-go-iam"
)
