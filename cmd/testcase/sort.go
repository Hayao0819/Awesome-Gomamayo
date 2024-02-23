package testcase

import (
	"sort"

	"github.com/spf13/cobra"
)

func sortCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "sort",
		Aliases: []string{"s"},
		Short:   "sort gomamayo.json",
		RunE: func(cmd *cobra.Command, args []string) error {

			var newTestcase Testcase = *testcase

			// sort gomamayo
			sort.SliceStable(newTestcase.Gomamayo, func(i, j int) bool {
				return newTestcase.Gomamayo[i].Word < newTestcase.Gomamayo[j].Word
			})

			// sort not-gomamayo
			sort.SliceStable(newTestcase.NotGomamayo, func(i, j int) bool {
				return newTestcase.NotGomamayo[i].Word < newTestcase.NotGomamayo[j].Word
			})

			// sort uncategorized
			sort.Strings(newTestcase.Uncategorized)

			// write
			if err := newTestcase.Write(testcasePath); err != nil {
				return err
			}

			return nil
		},
	}

	return &cmd
}
