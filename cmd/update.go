package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a password",
	Long:  `This subcommand updates a password`,
	Run:   updateHandler,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Not Implemented")
}
