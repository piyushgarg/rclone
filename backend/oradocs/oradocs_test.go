// Test Box filesystem interface
package oradocs_test

import (
	"testing"

	"github.com/rclone/rclone/backend/oradocs"
	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "qa:",
		NilObject:  (*oradocs.Object)(nil),
	})
}
