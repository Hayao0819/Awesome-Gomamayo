package testcase

import (
	"errors"

	"github.com/spf13/cobra"
)

func dupCommand() *cobra.Command {

	cmd := cobra.Command{
		Use:     "dup",
		Aliases: []string{"d"},
		Short:   "Check duplicate words",
		RunE: func(cmd *cobra.Command, args []string) error {

			errs := []error{}

			// check duplicate
			words := make(map[string]struct{})
			for _, g := range testcase.Gomamayo {
				if _, ok := words[g.Word]; ok {
					errs = append(errs, errors.New("duplicate word: "+g.Word))
				}
				words[g.Word] = struct{}{}
			}

			if len(errs) > 0 {
				for _, e := range errs {
					cmd.PrintErrln(e)
				}
				return errors.New("duplicate words")
			}

			return nil
		},
	}

	return &cmd
}
