package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "store a new password",
	Long:  `This subcommand stores a new password`,
	Run:   storeHandler,
}

func init() {
	rootCmd.AddCommand(storeCmd)
}

func storeHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Not Implemented")
}
