package main

import (
	"fmt"
	"io"

	"github.com/chengjingtao/jumper"
	"github.com/spf13/cobra"
)

type repoInspectCmd struct {
	name string

	out io.Writer
}

func newRepoInspectCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect [flags] [NAME]",
		Short: "Inspect a server repository",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				return fmt.Errorf("args format error, it should be add [NAME] [URL]")
			}

			inspect := &repoInspectCmd{out: out}
			inspect.name = args[0]

			return inspect.run()
		},
	}

	_ = cmd.Flags()

	return cmd
}

func (cmd *repoInspectCmd) run() error {
	areas, err := jumper.InspectRepo(cmd.name)
	if err != nil {
		return err
	}

	for _, area := range areas {
		cmd.out.Write([]byte(fmt.Sprintf("AREA: %s\n", area.Name)))
		for _, server := range area.Servers {
			cmd.out.Write([]byte(fmt.Sprintf(" => %s %s@%s %s\n", server.Name, server.User, server.IP, server.Description)))
		}
	}

	return nil
}
