package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomwright/kubo/internal"
	"github.com/tomwright/kubo/internal/config"
	"github.com/tomwright/kubo/internal/paths"
)

func update() *cobra.Command {
	var template string
	var overrides internal.OverrideFlag

	cmd := &cobra.Command{
		Use:   "update -t <template to use> -e <environment to use> <service name>",
		Short: "Update config files that are used during manifest generation.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, serviceName := range args {
				configPath := paths.ConfigFile(serviceName, envEnvironment, template)
				cfg, err := config.LoadFromFile(configPath)
				if err != nil {
					return fmt.Errorf("could not get config: %w", err)
				}

				for _, o := range overrides {
					if err := cfg.Set(o.Path, o.Value); err != nil {
						return fmt.Errorf("could not set property by key: %w", err)
					}
				}

				if err := config.SaveToFile(cfg, configPath); err != nil {
					return fmt.Errorf("could not save service config: %w", err)
				}

				fmt.Printf("%s [%s] updated\n", serviceName, envEnvironment)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&template, "template", "t", "service", "The template to use when generating manifests.")
	cmd.Flags().VarP(&overrides, "value", "v", "Set a specific config value. -v <path>=<value>")

	return cmd
}
