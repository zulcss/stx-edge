package cmd

import (
	"os"
	"github.com/spf13/cobra"
        "github.com/sirupsen/logrus"
	"github.com/zulcss/stx-edge/pkg/preflight"
)

var (
	Verbose 	bool
)

var rootCmd = &cobra.Command{
	Use: 	"stx-builder",
	Short:	"StarlingX Image Builder",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := InitLog(); err != nil {
			return err
		}

		preflight.PreflightCheck()
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(-1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func InitLog() error {
	if Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// general log output
	formatter := &logrus.TextFormatter{
                FullTimestamp:   true,
                TimestampFormat: "2006-01-02 15:04:05",
        }
        logrus.SetFormatter(formatter)

	return nil
}
