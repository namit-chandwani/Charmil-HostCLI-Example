package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// FooCmd represents the foo command
var FooCmd = &cobra.Command{
	Use:   "foo",
	Short: "Prints the message `foo called`",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("foo called")
	},
}

func init() {
}
