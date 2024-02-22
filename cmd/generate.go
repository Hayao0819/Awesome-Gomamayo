package cmd

import (
	"os"
	"os/exec"
	"path"

	"github.com/Hayao0819/awesome-gomamayo/gomamayo"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/nahi/tputils"
	"github.com/spf13/cobra"
)

func makeCmd() *cobra.Command {
	var pwd string

	cmd := cobra.Command{
		Use:     "make",
		Aliases: []string{"m"},
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

			tmplData := conf.ToTemplateData()

			data, err := tputils.ApplyTemplate(path.Join(pwd, "assets", "template.md"), tmplData)
			if err != nil {
				return err
			}

			if err := os.WriteFile(path.Join(pwd, "README.md"), data.Bytes(), 0644); err != nil {
				return err
			}

			exec.Command("pnpm", "markdownlint", "-f", path.Join(pwd, "README.md")).Run()

			return nil
		},
	}

	return &cmd
}

func init() {
	cobrautils.AddSubCmds(makeCmd())
}
