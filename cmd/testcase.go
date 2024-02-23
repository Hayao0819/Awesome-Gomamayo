package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func testCaseRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "testcase",
		Aliases: []string{"t"},
		Short:   "testcase",
	}

	cmd.AddCommand(testCaseSortCmd())
	return &cmd
}

func testCaseSortCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "sort",
		Aliases: []string{"s"},
		Short:   "sort gomamayo.json",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.PrintErrln("未実装")
			return nil
		},
	}

	return &cmd
}

func init() {
	cobrautils.AddSubCmds(testCaseRootCmd())
}
