package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var multiplyCmd = &cobra.Command{
	Use:     "multiply <a> <b>",
	Aliases: []string{"mul", "mult", "times"},
	Short:   "Multiply two numbers",
	Long:    "Carry out multiplication of two numbers.",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Multiply(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Printf("%s * %s = %s\n", args[0], args[1], result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
}
