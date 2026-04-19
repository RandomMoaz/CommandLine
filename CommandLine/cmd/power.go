package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:     "power <base> <exponent>",
	Aliases: []string{"pow", "exp"},
	Short:   "Raise a base to an exponent",
	Long:    "Compute base raised to the exponent (base ^ exponent).",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Power(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Printf("%s ^ %s = %s\n", args[0], args[1], result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)
}
