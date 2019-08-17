package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/chengjingtao/jumper"
	"github.com/spf13/cobra"
)

type toTargetCmd struct {
	name string
	out  io.Writer
}

func newToTargetCmd(out io.Writer) *cobra.Command {

	toCmd := &toTargetCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     "to [AreaName]/[targetServer] [FLAGS]",
		Short:   "ssh to target server",
		Example: `jumper to devops/int`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("you must specify a target server")
			}
			toCmd.name = args[0]
			return toCmd.Run()
		},
	}
	return cmd
}

func (cmd *toTargetCmd) Run() error {
	segs := strings.Split(cmd.name, "/")
	if len(segs) != 2 {
		return errors.New("ssh target should be areaName/targetServer")
	}

	areaName := segs[0]
	serverName := segs[1]
	config, err := jumper.GetConfig()
	if err != nil {
		return err
	}
	currentRepoName := config.Current.Repo
	if currentRepoName == "" {
		if len(config.Repos) == 0 {
			return errors.New("Empty repos found, execute repo add firstly")
		}
		currentRepoName = config.Repos[0].Name
	}

	// find target server
	var targetServer *jumper.Server
	areas, err := jumper.InspectRepo(currentRepoName)
	if err != nil {
		return err
	}
	for _, area := range areas {
		if area.Name != areaName {
			continue
		}
		for _, _server := range area.Servers {
			var server = _server
			if server.Name == serverName {
				targetServer = &server
			}
		}
	}

	if targetServer == nil {
		return fmt.Errorf("Not found server %s in area %s", serverName, areaName)
	}

	binary, err := exec.LookPath("ssh")
	if err != nil {
		return err
	}

	env := os.Environ()
	args := []string{
		"ssh",
		fmt.Sprintf("%s@%s", targetServer.User, targetServer.IP),
	}
	fmt.Printf("%s %s\n", binary, strings.Join(args, " "))
	err = syscall.Exec(binary, args, env)
	return err
}
