//go:build tools
// +build tools

// Copied from https://github.com/kubernetes/sample-controller/blob/master/hack/tools.go

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "k8s.io/code-generator"
)