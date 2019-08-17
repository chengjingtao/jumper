package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var version = "unkown"
var buildDate = "unkown"

func newVersionCmd(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%s(%s)\n", version, buildDate)
			return nil
		},
	}
	return cmd
}
