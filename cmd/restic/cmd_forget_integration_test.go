package main

import (
	"context"
	"testing"

	rtest "github.com/wutong-paas/restic/pkg/test"
)

func testRunForget(t testing.TB, gopts GlobalOptions, args ...string) {
	opts := ForgetOptions{}
	rtest.OK(t, runForget(context.TODO(), opts, gopts, args))
}
