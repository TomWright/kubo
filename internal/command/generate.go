package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomwright/kubo/internal"
	"github.com/tomwright/kubo/internal/config"
	"github.com/tomwright/kubo/internal/paths"
	"io/ioutil"
	"os"
)

func generate() *cobra.Command {
	var template string

	cmd := &cobra.Command{
		Use:   "generate -t <template to use> -e <environment to use> <service to generate>",
		Short: "Generate kubernetes manifests for the given service, template and environment.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, serviceName := range args {
				configPath := paths.ConfigFile(serviceName, envEnvironment, template)
				cfg, err := config.LoadFromFile(configPath)
				if err != nil {
					return fmt.Errorf("could not get config: %w", err)
				}

				files, err := ioutil.ReadDir(paths.TemplateDir(template))
				if err != nil {
					return fmt.Errorf("could not get list of template files: %w", err)
				}

				outputDir := paths.ManifestDir(serviceName, envEnvironment)
				if err := os.RemoveAll(outputDir); err != nil && !os.IsNotExist(err) {
					return fmt.Errorf("could not clear existing service config: %w", err)
				}

				if err := os.MkdirAll(outputDir, 0755); err != nil {
					return fmt.Errorf("could not create output directory: %w", err)
				}

				for _, f := range files {
					if err := internal.GenerateTemplateFile(f, serviceName, envEnvironment, template, cfg); err != nil {
						return fmt.Errorf("could not handle file: %w", err)
					}
				}
				fmt.Printf("%s [%s] generated\n", serviceName, envEnvironment)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&template, "template", "t", "service", "The template to use when generating manifests.")

	return cmd
}
