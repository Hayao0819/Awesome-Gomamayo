package testcase

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

var testcase *Testcase = nil
var testcasePath string = path.Join(pwd, "test", "gomamayo.json")
var pwd string

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "testcase",
		Aliases: []string{"t"},
		Short:   "testcase",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			pwd, err = os.Getwd()
			if err != nil {
				return err
			}

			testcase, err = GetTestCase(testcasePath)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(sortCmd())
	return &cmd
}
