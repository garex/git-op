package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

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
		out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		branch := strings.Replace(fmt.Sprintf("%s", out), "\n", "", -1)
		if len(args) > 0 {
			branch = args[0]
		}
		if "HEAD" == branch {
			log.Fatal("Unknown merge [NAME]")
		}
		if "master" == branch {
			return
		}

		out, err = exec.Command("git", "checkout", branch).CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		fmt.Printf("%s\n", out)

		out, err = exec.Command("git", "remote").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		if len(out) > 0 {
			out, err = exec.Command("git", "fetch", "origin", "master:master").CombinedOutput()
			if err != nil {
				log.Fatal(fmt.Sprintf("%s", out))
			}
			fmt.Printf("%s\n", out)
		}

		out, err = exec.Command("git", "show-ref", "--heads", "-s", "master").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		masterHash := out
		out, err = exec.Command("git", "merge-base", "master", branch).CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		mergeHash := out

		if !bytes.Equal(masterHash, mergeHash) {
			out, err = exec.Command("git", "rebase", "master").CombinedOutput()
			if err != nil {
				log.Fatal(fmt.Sprintf("%s", out))
			}
			fmt.Printf("%s\n", out)
		}

		out, err = exec.Command("git", "checkout", "master").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		fmt.Printf("%s\n", out)

		out, err = exec.Command("git", "merge", "--no-ff", branch).CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		fmt.Printf("%s\n", out)

		out, err = exec.Command("git", "branch", "-d", branch).CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		fmt.Printf("%s\n", out)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
