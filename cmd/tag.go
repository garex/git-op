package cmd

import (
	"github.com/spf13/cobra"
)

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:     "tag [VERSION]",
	Aliases: []string{"t"},
	Short:   "Release - Tag",
	Long: `'git op tag [VERSION]' creates version as a tag.

Assume we have last tag '1.2.3'.

* 'VERSION' is ommited or 'patch' passed -- creates next 'patch' version: '1.2.4'.
* 'minor' passed -- '1.3'.
* 'major' passed -- '2.0'.

Default behavior calls 'git op merge' before releasing.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
