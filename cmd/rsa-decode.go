package cmd

import (
	"encoder/encoders/rsa"
	"encoder/file"
	"errors"

	"github.com/spf13/cobra"
)

var key string

func init() {
	rsaDecodeCmd.Flags().StringVarP(&key, "key", "k", "", "key")

	rsaCmd.AddCommand(rsaDecodeCmd)
}

var rsaDecodeCmd = &cobra.Command{
	Use:   "decode [inputPath] [outputPath]",
	Short: "Decoding",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if key == "" {
			panic(errors.New("Key is required"))
		}

		inputPath, outputPath := args[0], args[1]

		encodedText, err := file.Read(inputPath)

		if err != nil {
			panic(err)
		}

		decodedText, err := rsa.Decode(encodedText, key)

		if err != nil {
			panic(err)
		}

		if err = file.Write(outputPath, decodedText); err != nil {
			panic(err)
		}
	},
}
