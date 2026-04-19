package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var modCmd = &cobra.Command{
	Use:     "mod <a> <b>",
	Aliases: []string{"modulo", "remainder"},
	Short:   "Compute the remainder of a / b",
	Long:    "Compute a mod b (the remainder after dividing a by b). Accepts floats.",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Mod(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Printf("%s mod %s = %s\n", args[0], args[1], result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(modCmd)
}
