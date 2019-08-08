package cmd

import (
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

	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
