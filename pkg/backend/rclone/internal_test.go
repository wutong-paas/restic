package rclone

import (
	"context"
	"os/exec"
	"testing"

	"github.com/wutong-paas/restic/pkg/errors"
	"github.com/wutong-paas/restic/pkg/restic"
	rtest "github.com/wutong-paas/restic/pkg/test"
)

// restic should detect rclone exiting.
func TestRcloneExit(t *testing.T) {
	dir := rtest.TempDir(t)
	cfg := NewConfig()
	cfg.Remote = dir
	be, err := Open(context.TODO(), cfg, nil)
	var e *exec.Error
	if errors.As(err, &e) && e.Err == exec.ErrNotFound {
		t.Skipf("program %q not found", e.Name)
		return
	}
	rtest.OK(t, err)
	defer func() {
		// ignore the error as the test will kill rclone (see below)
		_ = be.Close()
	}()

	err = be.cmd.Process.Kill()
	rtest.OK(t, err)
	t.Log("killed rclone")

	for i := 0; i < 10; i++ {
		_, err = be.Stat(context.TODO(), restic.Handle{
			Name: "foo",
			Type: restic.PackFile,
		})
		rtest.Assert(t, err != nil, "expected an error")
	}
}

// restic should detect rclone startup failures
func TestRcloneFailedStart(t *testing.T) {
	cfg := NewConfig()
	// exits with exit code 1
	cfg.Program = "false"
	_, err := Open(context.TODO(), cfg, nil)
	var e *exec.ExitError
	if !errors.As(err, &e) {
		// unexpected error
		rtest.OK(t, err)
	}
}
