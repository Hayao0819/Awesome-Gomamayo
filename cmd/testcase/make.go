package testcase

import (
	"path"

	"github.com/spf13/cobra"
)

func makeCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "make",
		Aliases: []string{"m"},
		Short:   "make gomamayo.json",
		RunE: func(cmd *cobra.Command, args []string) error {
			// make gomamayo.txt
			txtPath := path.Join(pwd, "test", "gomamayo.txt")
			csvPath := path.Join(pwd, "test", "gomamayo.csv")
			jsonPath := path.Join(pwd, "test", "gomamayo.json")

			// write
			if err := testcase.Write(jsonPath, csvPath, txtPath); err != nil {
				return err
			}

			return nil
		},
	}

	return &cmd
}
