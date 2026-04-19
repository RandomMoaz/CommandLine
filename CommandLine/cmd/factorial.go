package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var factorialCmd = &cobra.Command{
	Use:     "factorial <n>",
	Aliases: []string{"fact", "fac"},
	Short:   "Compute n! (factorial)",
	Long:    "Compute the factorial of a non-negative integer. Max supported input is 170.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Factorial(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("%s! = %s\n", args[0], result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(factorialCmd)
}
