package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var calcCmd = &cobra.Command{
	Use:     "calc <expression>",
	Aliases: []string{"eval", "compute"},
	Short:   "Evaluate a full arithmetic expression",
	Long: "Evaluate an arithmetic expression with operator precedence and parentheses.\n" +
		"Supported operators: + - * / % ^ and unary minus. Numbers can be integers,\n" +
		"decimals, or use scientific notation (e.g. 1.5e3).\n\n" +
		"Examples:\n" +
		"  zero calc \"2 + 3 * 4\"        -> 14\n" +
		"  zero calc \"(2 + 3) * 4\"      -> 20\n" +
		"  zero calc \"2 ^ 10\"           -> 1024\n" +
		"  zero calc \"-5 + 10 / 2\"      -> 0\n",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Join args so both `calc "2 + 3"` and `calc 2 + 3` work.
		expr := strings.Join(args, " ")
		result, err := Evaluate(expr)
		if err != nil {
			return err
		}
		fmt.Printf("%s = %s\n", expr, result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)
}
