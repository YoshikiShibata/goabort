// Copyright © 2020 Yoshiki Shibata. All rights reserved.

package goabort

import (
	"log"
	"runtime"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type caller struct {
	file string
	line int
}

var lock sync.Mutex
var calledMap = make(map[caller]bool)

// Once returns codes.Abort error when this function is called first.
func Once(err error) error {
	lock.Lock()
	defer lock.Unlock()

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Cannot call runtime.Caller(1)")
	}

	k := caller{file: file, line: line}
	if calledMap[k] {
		return err
	}

	calledMap[k] = true
	log.Printf("%s:%d Aborted return", file, line)
	return status.Error(codes.Aborted, "replacement")
}
