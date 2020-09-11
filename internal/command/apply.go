package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomwright/kubo/internal/paths"
	"os/exec"
)

func apply() *cobra.Command {
	var environment string

	cmd := &cobra.Command{
		Use:   "apply -e <environment to use> <service to generate>",
		Short: "Apply kubernetes manifests for the given service + environment.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, serviceName := range args {
				manifestDir := paths.ManifestDir(serviceName, environment)

				fmt.Printf("Applying %s [%s]\n", serviceName, environment)

				cmd := exec.Command(
					"kubectl",
					"apply",
					"-f",
					manifestDir,
				)

				fmt.Printf("%s\n", cmd.String())

				output, err := cmd.CombinedOutput()
				if err != nil {
					return fmt.Errorf("command finished with error: %w: %s", err, string(output))
				}

				fmt.Print(string(output))

				fmt.Printf("%s [%s] applied\n", serviceName, environment)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&environment, "environment", "e", "default", "The environment we're applying manifests for.")

	return cmd
}
