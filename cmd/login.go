package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to use this service",
	Long:  `This subcommand logs you in`,
	Run:   loginHandler,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func loginHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Not Implemented")
}
