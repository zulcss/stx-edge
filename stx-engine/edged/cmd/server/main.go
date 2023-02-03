package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:	"edged",
	Short:	"starlingx edge api server",
	Run: func(cmd *cobra.Command, args []string) {},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed executing root command")
	}
}

func init() {}
