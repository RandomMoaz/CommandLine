package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add <a> <b>",
	Aliases: []string{"addition", "plus"},
	Short:   "Add two numbers",
	Long:    "Carry out addition of two numbers.",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Add(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Printf("%s + %s = %s\n", args[0], args[1], result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
