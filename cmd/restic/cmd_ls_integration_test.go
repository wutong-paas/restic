package main

import (
	"context"
	"strings"
	"testing"

	rtest "github.com/wutong-paas/restic/pkg/test"
)

func testRunLs(t testing.TB, gopts GlobalOptions, snapshotID string) []string {
	buf, err := withCaptureStdout(func() error {
		gopts.Quiet = true
		opts := LsOptions{}
		return runLs(context.TODO(), opts, gopts, []string{snapshotID})
	})
	rtest.OK(t, err)
	return strings.Split(buf.String(), "\n")
}
