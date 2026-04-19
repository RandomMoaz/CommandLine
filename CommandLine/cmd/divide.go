package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var divideCmd = &cobra.Command{
	Use:     "divide <a> <b>",
	Aliases: []string{"div"},
	Short:   "Divide the first number by the second",
	Long:    "Carry out division: a / b. Division by zero returns an error.",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Divide(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Printf("%s / %s = %s\n", args[0], args[1], result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(divideCmd)
}
