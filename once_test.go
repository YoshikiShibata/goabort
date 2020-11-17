// Copyright © 2020 Yoshiki Shibata. All rights reserved.

package goabort_test

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/YoshikiShibata/goabort"
)

func TestOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := errors.New("own error")
		newErr := goabort.Once(err)
		if i == 0 {
			if status.Code(newErr) != codes.Aborted {
				t.Fatalf("err is %v, want %v", newErr, codes.Aborted)
			}
		} else if newErr != err {
			t.Errorf("err is not identical")
		}
	}
}
