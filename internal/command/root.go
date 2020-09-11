package command

import (
	"github.com/spf13/cobra"
	"github.com/tomwright/kubo/internal/paths"
)

var (
	envBasePath    string
	envEnvironment string
)

var RootCMD = &cobra.Command{
	Use:        "kubo",
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "A small helper to manage kubernetes configurations.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if envBasePath != "" {
			paths.StdPath.SetBase(envBasePath)
		}
	},
}

func init() {
	RootCMD.AddCommand(
		generate(),
		update(),
		apply(),
	)

	RootCMD.PersistentFlags().StringVarP(&envBasePath, "base", "b", ".", "Full path to kubo base directory.")
	RootCMD.PersistentFlags().StringVarP(&envEnvironment, "environment", "e", "default", "The environment to work with.")
}
