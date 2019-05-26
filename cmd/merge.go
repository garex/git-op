package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:     "merge [NAME]",
	Aliases: []string{"m"},
	Short:   "Finish - Merge",
	Long: `'git op merge [NAME]' merges branch 'NAME' into 'master'.

When 'NAME' is ommitted, current branch is used. On 'master' branch it do nothing.
	
Default behavior pulls latest 'master' branch from 'origin' remote and rebases onto it.
Tags are fetched too.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("merge called")
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
