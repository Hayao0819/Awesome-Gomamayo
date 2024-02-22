package cmd

import (
	"errors"
	"net/url"
	"os"
	"path"

	"github.com/Hayao0819/awesome-gomamayo/gomamayo"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func checkCmd() *cobra.Command {
	var pwd string
	cmd := cobra.Command{
		Use:     "check",
		Aliases: []string{"c"},
		Short:   "check gomamayo.toml",
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

			var urls []string
			var errs []error

			for _, r := range conf.Repos {
				urls = append(urls, r.Url)
			}
			for _, w := range conf.Websites {
				urls = append(urls, w.Url)
			}

			for _, u := range urls {
				if _, err := url.ParseRequestURI(u); err != nil {
					errs = append(errs, err)
				}
			}

			if len(errs) > 0 {
				return errors.New("invalid url")
			}

			return nil

		},
	}

	return &cmd
}

func init() {
	cobrautils.AddSubCmds(checkCmd())
}
