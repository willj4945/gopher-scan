package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goscan",
	Short: "A fast, concurrent network and port scanner",
	Long:  `go-scan — comprehensive network scanner for asset discovery and port enumeration.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
