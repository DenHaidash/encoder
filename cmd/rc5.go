package cmd

import (
	"github.com/spf13/cobra"
)

var rc5key string

func init() {
	rootCmd.AddCommand(rc5Cmd)
}

var rc5Cmd = &cobra.Command{
	Use:   "rc5",
	Short: "RC5 algorithm",
}
