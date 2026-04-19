package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "subtract <a> <b>",
	Aliases: []string{"sub", "minus"},
	Short:   "Subtract the second number from the first",
	Long:    "Carry out subtraction: a - b.",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Subtract(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Printf("%s - %s = %s\n", args[0], args[1], result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
