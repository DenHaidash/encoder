package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rsaCmd)
}

var rsaCmd = &cobra.Command{
	Use:   "rsa",
	Short: "RSA algorithm",
}
