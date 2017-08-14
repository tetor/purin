package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

		path := "/Users/tetor/purin/src/github.com/tetor/purin"
		url := "https://github.com/tetor/purin.git"

		_, err := git.PlainClone(path, false, &git.CloneOptions{
			URL: url,
			Progress: os.Stdout,
		})
		if err != nil { fmt.Errorf("error %v", err) }


		confComd := exec.Command(fmt.Sprintf("./configure --prefix=/Users/tetor/purin"))
		makeComd := exec.Command("make")
		makeIComd := exec.Command("make install")

		confComd.Dir = path
		makeComd.Dir = path
		makeIComd.Dir = path

		out1, _ := confComd.Output()
		fmt.Printf("%s", out1)

		out2, _ := makeComd.Output()
		fmt.Printf("make %s", out2)

		out3, _ := makeIComd.Output()
		fmt.Printf("make in %s", out3)
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
