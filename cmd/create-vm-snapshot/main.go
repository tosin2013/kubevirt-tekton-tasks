package main

import (
	"net/http"

	goarg "github.com/alexflint/go-arg"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/exit"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/log"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/output"
	res "github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/results"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/zerrors"
	. "github.com/tosin2013/kubevirt-tekton-tasks/modules/create-vm-snapshot/pkg/constants"
	createvmsnapshot "github.com/tosin2013/kubevirt-tekton-tasks/modules/create-vm-snapshot/pkg/snapshots"
	"github.com/tosin2013/kubevirt-tekton-tasks/modules/create-vm-snapshot/pkg/utils/parse"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	defer exit.HandleExit()

	cliOptions := &parse.CLIOptions{}
	goarg.MustParse(cliOptions)

	// Convert string to zapcore.Level
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(cliOptions.GetDebugLevel())); err != nil {
		exit.ExitOrDieFromError(InvalidCLIInputExitCode, err)
	}

	// Initialize logger with the specified debug level
	logger := log.InitLogger(logLevel)
	defer logger.Sync()

	// Initialize CLI options
	if err := cliOptions.Init(); err != nil {
		exit.ExitOrDieFromError(InvalidCLIInputExitCode, err)
	}

	log.Logger().Debug("parsed arguments", zap.Reflect("cliOptions", cliOptions))

	// Create a new SnapshotManager
	snapshotManager, err := createvmsnapshot.NewSnapshotManager(cliOptions)
	if err != nil {
		exit.ExitOrDieFromError(SnapshotManagerErrorCode, err)
	}

	// Create a new snapshot for the specified VM
	newSnapshot, err := snapshotManager.CreateSnapshot()
	if err != nil {
		exit.ExitOrDieFromError(CreateSnapshotErrorCode, err,
			zerrors.IsStatusError(err, http.StatusNotFound, http.StatusConflict, http.StatusUnprocessableEntity),
		)
	}

	// Record and log the results
	results := map[string]string{
		"snapshotName": newSnapshot.Name,
		"vmName":       newSnapshot.VMName,
		"creationTime": newSnapshot.CreationTime,
	}

	log.Logger().Debug("recording results", zap.Reflect("results", results))
	if err := res.RecordResults(results); err != nil {
		exit.ExitOrDieFromError(WriteResultsExitCode, err)
	}

	// Pretty print the snapshot details
	output.PrettyPrint(newSnapshot, output.OutputType(cliOptions.Output))
}

