package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sqrtCmd = &cobra.Command{
	Use:     "sqrt <n>",
	Aliases: []string{"squareroot"},
	Short:   "Compute the square root of a number",
	Long:    "Compute the (non-negative) square root of a non-negative number.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Sqrt(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("sqrt(%s) = %s\n", args[0], result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sqrtCmd)
}
