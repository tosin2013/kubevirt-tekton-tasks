package parse

type CLIOptions struct {
    DebugLevel string `arg:"-d,--debug" help:"Debug level"`
    VMName     string `arg:"-v,--vm-name" help:"Name of the VM to snapshot" required:"true"`
    Output     string `arg:"-o,--output" help:"Output format (e.g., json, yaml)" default:"json"`
}

func (cli *CLIOptions) GetDebugLevel() string {
    return cli.DebugLevel
}

func (cli *CLIOptions) Init() error {
    // Additional initialization if needed
    return nil
}
