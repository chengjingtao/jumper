package main

import (
	"fmt"
	"io"

	"github.com/chengjingtao/jumper"
	"github.com/spf13/cobra"
)

type repoAddCmd struct {
	name     string
	url      string
	username string
	password string

	out io.Writer
}

func newRepoAddCmd(out io.Writer) *cobra.Command {
	add := &repoAddCmd{out: out}

	cmd := &cobra.Command{
		Use:   "add [flags] [NAME] [URL]",
		Short: "Add a server repository",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 2 {
				return fmt.Errorf("args format error, it should be add [NAME] [URL]")
			}

			add.name = args[0]
			add.url = args[1]

			return add.run()
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&add.password, "password", "p", "", "password of the repository")
	flags.StringVarP(&add.username, "user", "u", "", "user of the repository")

	return cmd
}

func (cmd *repoAddCmd) run() error {
	repo := jumper.Repo{
		Name: cmd.name,
		Url:  cmd.url,
	}
	repo.SetUsernamePassword(cmd.username, cmd.password)
	return jumper.AppendRepo(repo)
}
