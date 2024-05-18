package snapshots

import (
	"testing"

	"github.com/tosin2013/kubevirt-tekton-tasks/modules/create-vm-snapshot/pkg/utils/parse"
)

func TestCreateSnapshot(t *testing.T) {
    cliOptions := &parse.CLIOptions{VMName: "test-vm"}
    snapshotManager, err := NewSnapshotManager(cliOptions)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    snapshot, err := snapshotManager.CreateSnapshot()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if snapshot.VMName != "test-vm" {
        t.Errorf("expected VMName 'test-vm', got %v", snapshot.VMName)
    }
}
