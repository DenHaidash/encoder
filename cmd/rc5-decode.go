package cmd

import (
	"encoder/encoders/rc5"
	"encoder/file"
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	rc5DecodeCmd.Flags().StringVarP(&rc5key, "key", "k", "", "key")

	rc5Cmd.AddCommand(rc5DecodeCmd)
}

var rc5DecodeCmd = &cobra.Command{
	Use:   "decode [inputPath] [outputPath]",
	Short: "Decoding",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if rc5key == "" {
			panic(errors.New("Key is required"))
		}

		inputPath, outputPath := args[0], args[1]

		encodedText, err := file.Read(inputPath)

		if err != nil {
			panic(err)
		}

		decodedText, err := rc5.Decode(encodedText, rc5key)

		if err != nil {
			panic(err)
		}

		if err = file.Write(outputPath, decodedText); err != nil {
			panic(err)
		}
	},
}
