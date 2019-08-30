package main

import (
	"fmt"
	"io"

	"github.com/chengjingtao/jumper"
	"github.com/spf13/cobra"
)

type repoListCmd struct {
	out io.Writer
}

func newRepoListCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [flags] [NAME]",
		Short: "List a server repository",
		RunE: func(cmd *cobra.Command, args []string) error {

			list := &repoListCmd{out: out}

			return list.run()
		},
	}

	_ = cmd.Flags()

	return cmd
}

func (cmd *repoListCmd) run() error {
	repos, err := jumper.ListRepo()
	if err != nil {
		return err
	}

	if len(repos) == 0 {
		_, _ = cmd.out.Write([]byte("Empty repos ~\n"))
	}

	for _, item := range repos {
		_, _ = cmd.out.Write([]byte(fmt.Sprintf("%s  %s\n", item.Name, item.Url)))
	}

	return nil
}
