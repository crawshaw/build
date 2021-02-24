// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maintpb

// Run "go generate" in this directory to update. You need to install gRPC-Go:
//
//	go get google.golang.org/protobuf/cmd/protoc-gen-go \
//	  google.golang.org/grpc/cmd/protoc-gen-go-grpc

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative maintner.proto
