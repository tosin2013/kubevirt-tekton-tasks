package parse

type CLIOptions struct {
    VMName     string
    Namespace  string
    DebugLevel string
    Output     string  // Add this line if Output is supposed to be a field
}

func (cli *CLIOptions) GetDebugLevel() string {
    return cli.DebugLevel
}

func (cli *CLIOptions) Init() error {
    // Additional initialization if needed
    return nil
}
