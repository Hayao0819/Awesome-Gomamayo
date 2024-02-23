package testcase

import "github.com/spf13/cobra"

func checkCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "check",
		Aliases: []string{"c"},
		Short:   "check gomamayo.json",
		RunE: func(cmd *cobra.Command, args []string) error {

			// check duplication
			dupCmd := dupCommand()
			if err := dupCmd.Execute(); err != nil {
				return err
			}

			return nil
		},
	}

	return &cmd

}
