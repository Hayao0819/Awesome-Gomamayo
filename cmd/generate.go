package cmd

import (
	"os"
	"path"

	"github.com/Hayao0819/awesome-gomamayo/gomamayo"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)




func generateCmd() *cobra.Command {
	var pwd string

	cmd := cobra.Command{
		Use:     "generate",
		Aliases: []string{"g", "gen"},
		Short:   "generate README.md",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			pwd, err = os.Getwd()
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := gomamayo.Read(path.Join(pwd, "gomamayo.toml"))
			if err != nil {
				return err
			}

			cmd.Println(*conf)
			cmd.Println(conf.ToTemplateData())
			return nil
		},
	}

	return &cmd
}

func init() {
	cobrautils.AddSubCmds(generateCmd())
}
