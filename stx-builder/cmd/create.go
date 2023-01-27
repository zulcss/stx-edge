package cmd

import "github.com/spf13/cobra"

var createCommand = &cobra.Command{
	Use:	"create",
	Short:	"create starlingx image artifcat",
}

func init() {
	rootCmd.AddCommand(createCommand)
}
