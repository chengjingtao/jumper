package main

import (
	"io"

	"github.com/spf13/cobra"
)

func newRepoCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo [FLAGS] add|remove|list|update|inspect [ARGS]",
		Short: "Add, remove list update, and inspect repository",
		Long:  "Add, remove list update, and inspect repository",
	}

	cmd.AddCommand(newRepoAddCmd(out))
	cmd.AddCommand(newRepoUpdateCmd(out))
	cmd.AddCommand(newRepoInspectCmd(out))
	cmd.AddCommand(newRepoRemoveCmd(out))
	cmd.AddCommand(newRepoListCmd(out))

	return cmd
}
