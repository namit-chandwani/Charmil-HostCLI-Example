package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BarCmd represents the bar command
var BarCmd = &cobra.Command{
	Use:   "bar",
	Short: "Prints the message `bar called`",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bar called")
	},
}

func init() {
}
