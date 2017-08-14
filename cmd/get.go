package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
)

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get wild",
	Long: `Desc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")

		_, err := git.PlainClone("/Users/tetor/purin/src/purin", false, &git.CloneOptions{
			URL: "https://github.com/tetor/purin.git",
			Progress: os.Stdout,
		})

		if err != nil {
			fmt.Errorf("error %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(GetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
