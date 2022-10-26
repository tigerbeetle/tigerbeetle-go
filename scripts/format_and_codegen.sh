#!/usr/bin/env bash

set -eu

GO_IMAGE=golang:1.19

docker run -v $(pwd):/wrk "$GO_IMAGE" bash -c "cd /wrk && gofmt -w -s ."

docker run -v $(pwd):/wrk "$GO_IMAGE" bash -c "go install golang.org/x/tools/cmd/stringer@latest && go generate /wrk/pkg/types/main.go"
