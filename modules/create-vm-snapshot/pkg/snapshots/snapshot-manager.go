package snapshots

import (
	"fmt"

	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/zerrors"
	"github.com/tosin2013/kubevirt-tekton-tasks/modules/create-vm-snapshot/pkg/utils/parse"
)

type SnapshotManager struct {
    Options *parse.CLIOptions
}

func NewSnapshotManager(options *parse.CLIOptions) (*SnapshotManager, error) {
    return &SnapshotManager{Options: options}, nil
}

type Snapshot struct {
    Name         string
    VMName       string
    CreationTime string
}

func (sm *SnapshotManager) CreateSnapshot() (*Snapshot, error) {
    if sm.Options.VMName == "" {
        return nil, zerrors.New(fmt.Sprintf("VM name is required"))
    }

    snapshot := &Snapshot{
        Name:         "snapshot-" + sm.Options.VMName,
        VMName:       sm.Options.VMName,
        CreationTime: "2024-05-18T12:34:56Z",
    }

    return snapshot, nil
}
