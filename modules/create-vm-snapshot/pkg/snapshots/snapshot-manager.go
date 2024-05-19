package snapshots

import (
	"github.com/tosin2013/kubevirt-tekton-tasks/modules/create-vm-snapshot/pkg/utils/parse"
	"github.com/tosin2013/kubevirt-tekton-tasks/modules/shared/pkg/zerrors"
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
        return nil, zerrors.New("VM name is required")
    }

    snapshot := &Snapshot{
        Name:         "snapshot-" + sm.Options.VMName,
        VMName:       sm.Options.VMName,
        CreationTime: "2024-05-18T12:34:56Z",
    }

    return snapshot, nil
}
