package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:     "push",
	Aliases: []string{"p"},
	Short:   "Push latest changes in 'master' to remote 'origin'",
	Long: `'git op push' push latest changes in 'master' with tags

Default behavior calls 'git op tag' before pushin, which will create patch version bump.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tagCmd.Run(cmd, []string{})

		out, err := exec.Command("git", "push", "origin", "master").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		fmt.Printf("%s\n", out)

		out, err = exec.Command("git", "push", "origin", "--tags").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		fmt.Printf("%s\n", out)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
