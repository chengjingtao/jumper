package main

import (
	"fmt"
	"io"

	"github.com/chengjingtao/jumper"
	"github.com/spf13/cobra"
)

type repoRemoveCmd struct {
	name string

	out io.Writer
}

func newRepoRemoveCmd(out io.Writer) *cobra.Command {
	remove := &repoRemoveCmd{out: out}

	cmd := &cobra.Command{
		Use:   "remove [flags] [NAME] [URL]",
		Short: "Remove a server repository",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return fmt.Errorf("args format error, it should be remove [NAME]")
			}

			remove.name = args[0]

			return remove.run()
		},
	}

	_ = cmd.Flags()

	return cmd
}

func (cmd *repoRemoveCmd) run() error {
	return jumper.RemoveRepo(cmd.name)
}
