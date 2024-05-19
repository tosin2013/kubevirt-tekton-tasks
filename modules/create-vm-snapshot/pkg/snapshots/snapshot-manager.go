package snapshots

import (
	"context"
	"time"

	"github.com/tosin2013/kubevirt-tekton-tasks/modules/create-vm-snapshot/pkg/utils/parse"
	"github.com/tosin2013/kubevirt-tekton-tasks/modules/shared/pkg/zerrors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	kubevirtv1 "kubevirt.io/api/snapshot/v1alpha1"
	"kubevirt.io/client-go/kubecli"
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

	// Initialize the KubeVirt client
	kubevirtClient, err := kubecli.GetKubevirtClient()
	if err != nil {
		return nil, err
	}

	// Define the snapshot
	snapshot := &kubevirtv1.VirtualMachineSnapshot{
		ObjectMeta: metav1.ObjectMeta{
			Name: "snapshot-" + sm.Options.VMName,
		},
		Spec: kubevirtv1.VirtualMachineSnapshotSpec{
			Source: corev1.TypedLocalObjectReference{
				APIGroup: pointer.StringPtr("kubevirt.io"),
				Kind:     "VirtualMachine",
				Name:     sm.Options.VMName,
			},
		},
	}

	// Create the snapshot
	createdSnapshot, err := kubevirtClient.VirtualMachineSnapshot(sm.Options.Namespace).Create(context.Background(), snapshot, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return &Snapshot{
		Name:         createdSnapshot.Name,
		VMName:       sm.Options.VMName,
		CreationTime: createdSnapshot.CreationTimestamp.Time.Format(time.RFC3339),
	}, nil
}
