package parse

import (
	"testing"

	"github.com/alexflint/go-arg"
)

func TestCLIOptions(t *testing.T) {
    args := []string{"--vm-name", "test-vm", "--debug", "1"}
    cliOptions := &CLIOptions{}
    arg.MustParse(cliOptions)

    if cliOptions.VMName != "test-vm" {
        t.Errorf("expected VMName 'test-vm', got %v", cliOptions.VMName)
    }
}
