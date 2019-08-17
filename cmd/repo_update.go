package main

import (
	"fmt"
	"io"

	"github.com/chengjingtao/jumper"
	"github.com/spf13/cobra"
)

type repoUpdateCmd struct {
	name string

	out io.Writer
}

func newRepoUpdateCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [flags] [NAME]",
		Short: "Update a server repository",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				return fmt.Errorf("args format error, it should be add [NAME] [URL]")
			}

			update := &repoUpdateCmd{out: out}
			update.name = args[0]

			return update.run()
		},
	}

	_ = cmd.Flags()

	return cmd
}

func (cmd *repoUpdateCmd) run() error {
	return jumper.UpdateRepo(cmd.name, cmd.out)
}
