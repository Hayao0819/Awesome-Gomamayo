package cmd

import (
	"os"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

var pwd string

func Root() *cobra.Command {
	cmd := cobra.Command{
		Use:   "mamayo",
		Short: "mamayo is a tool to manage awesome Gomamayo",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			pwd, err = os.Getwd()
			if err != nil {
				return err

			}
			return nil
		},
		SilenceUsage: true,
	}

	cobrautils.AddSubCmdsToRoot(&cmd)
	cmd.CompletionOptions.DisableDefaultCmd = true

	return &cmd
}
