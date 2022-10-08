package cmd

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2022
	Description: Apx is a wrapper around apt to make it works inside a container
	from outside, directly on the host.
*/

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/apx/core"
)

func updateUsage(*cobra.Command) error {
	fmt.Print(`Description: 
Update list of available packages.

Usage:
  apx update

Examples:
  apx update
`)
	return nil
}

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update list of available packages.",
		RunE:  update,
	}
	cmd.SetUsageFunc(updateUsage)
	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	sys := cmd.Flag("sys").Value.String() == "true"
	command := append([]string{}, core.GetPkgManager(sys)...)
	command = append(command, "update")
	command = append(command, args...)

	if sys {
		core.AlmostRun(command...)
		return nil
	}

	core.RunContainer(command...)
	return nil
}