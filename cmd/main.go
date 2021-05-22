package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spreedly",
	Short: "Interact with the Spreedly API",
	Long:  `Spreedly is a CLI library for interfacing with your environments programatically.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(gatewayCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "0.1",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
