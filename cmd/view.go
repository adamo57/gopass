package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view all of your passwords",
	Long:  `This subcommand views all of your passwords`,
	Run:   viewHandler,
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

func viewHandler(cmd *cobra.Command, args []string) {
	fmt.Println("hello called")
}
