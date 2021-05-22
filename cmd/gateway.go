package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Interact with user created Spreedly gateways",
	Long:  `Available options: get, list.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Token:", args[0])
	},
}
