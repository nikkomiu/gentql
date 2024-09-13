package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

type Option interface {
	CmdOpt(*cobra.Command)
}

type WithOutputOption struct {
	stdout io.Writer
	stderr io.Writer
}

func WithOutput(stdout, stderr io.Writer) WithOutputOption {
	return WithOutputOption{stdout: stdout, stderr: stderr}
}

func (o WithOutputOption) CmdOpt(cmd *cobra.Command) {
	cmd.SetOut(o.stdout)
	cmd.SetErr(o.stderr)
}

type WithArgsOption struct {
	args []string
}

func WithArgs(args []string) WithArgsOption {
	return WithArgsOption{args: args}
}

func (o WithArgsOption) CmdOpt(cmd *cobra.Command) {
	cmd.SetArgs(o.args)
}
