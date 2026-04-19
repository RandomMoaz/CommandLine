package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is set at build time with:
//
//   go build -ldflags "-X go-cl/cmd.Version=v0.2.0" .
//
// It defaults to "dev" for local builds.
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the zero CLI version",
	Long:  "Print the version string of the zero CLI. Set at build time via -ldflags.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("zero %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
