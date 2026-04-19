package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zero",
	Short: "zero is a cli tool for performing basic mathematical operations",
	Long: "zero is a cli tool for performing basic mathematical operations:\n" +
		"  add, subtract, multiply, divide, power, sqrt, mod, factorial, calc.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command. Errors from subcommands are printed here
// and the process exits with a non-zero status so shell pipelines behave.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
