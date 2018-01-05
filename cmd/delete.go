package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete one or many passwords",
	Long:  `This subcommand deletes one or many of your passwords`,
	Run:   deleteHandler,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Not Implemented")
}
