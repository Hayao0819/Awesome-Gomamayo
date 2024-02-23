package cmd

import (
	"errors"
	"net/url"
	"path"

	"github.com/Hayao0819/awesome-gomamayo/gomamayo"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func urlCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "check URL in gomamayo.toml",

		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := gomamayo.Read(path.Join(pwd, "gomamayo.toml"))
			if err != nil {
				return err
			}

			var urls []string
			var errs []error

			for _, r := range conf.Codes {
				urls = append(urls, r.Url)
			}
			for _, w := range conf.Websites {
				urls = append(urls, w.Url)
			}

			for _, u := range urls {
				cmd.PrintErrln("checking", u)
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
	cobrautils.AddSubCmds(urlCmd())
}
