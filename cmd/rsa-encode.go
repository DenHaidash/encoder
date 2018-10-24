package cmd

import (
	"encoder/encoders/rsa"
	"encoder/file"

	"github.com/spf13/cobra"
)

func init() {
	rsaCmd.AddCommand(rsaEncodeCmd)
}

var rsaEncodeCmd = &cobra.Command{
	Use:   "encode [inputPath] [outputPath]",
	Short: "Encoding",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, outputPath := args[0], args[1]

		plainText, err := file.Read(inputPath)

		if err != nil {
			panic(err)
		}

		encodedText, keys, err := rsa.Encode(plainText)

		if err != nil {
			panic(err)
		}

		println("Public key: ", keys.PublicKey)
		println("Private key: ", keys.PrivateKey)

		if err = file.Write(outputPath, encodedText); err != nil {
			panic(err)
		}
	},
}
