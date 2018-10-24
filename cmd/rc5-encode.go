package cmd

import (
	"encoder/encoders/rc5"
	"encoder/file"
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	rc5EncodeCmd.Flags().StringVarP(&rc5key, "key", "k", "", "key")

	rc5Cmd.AddCommand(rc5EncodeCmd)
}

var rc5EncodeCmd = &cobra.Command{
	Use:   "encode [inputPath] [outputPath]",
	Short: "Encoding",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if rc5key == "" {
			panic(errors.New("Key is required"))
		}

		inputPath, outputPath := args[0], args[1]

		plainText, err := file.Read(inputPath)

		if err != nil {
			panic(err)
		}

		encodedText := rc5.Encode(plainText, rc5key)

		if err = file.Write(outputPath, encodedText); err != nil {
			panic(err)
		}
	},
}
