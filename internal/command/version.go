package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomwright/kubo/internal"
)

func version() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the kubo version.",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(internal.Version)
		},
	}

	return cmd
}
