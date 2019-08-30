package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := newRootCMD(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func newRootCMD(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "jump",
		Short:             "jump to target server",
		Long:              "jump to target server",
		SilenceUsage:      true,
		PersistentPreRun:  func(cmd *cobra.Command, args []string) {},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {},
	}
	flags := cmd.PersistentFlags()
	out := cmd.OutOrStdout()
	cmd.AddCommand(newRepoCmd(out))
	cmd.AddCommand(newToTargetCmd(out))
	cmd.AddCommand(newVersionCmd(out))

	err := flags.Parse(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cmd
}
