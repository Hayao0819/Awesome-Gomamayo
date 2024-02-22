package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := cobra.Command{
		Use:           "mamayo",
		Short:         "mamayo is a tool to manage awesome Gomamayo",
		SilenceUsage:  true,
	}

	cobrautils.AddSubCmdsToRoot(&cmd)

	return &cmd
}
