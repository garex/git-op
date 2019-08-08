package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
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
* 'minor' passed -- '1.3.0'.
* 'major' passed -- '2.0.0'.

Default behavior calls 'git op merge' before releasing.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mergeCmd.Run(cmd, []string{})

		out, err := exec.Command("git", "tag", "--list", "--points-at", "master").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		if len(out) > 0 {
			log.Fatal(fmt.Sprintf("Master already tagged: %s", out))
		}

		version := "patch"
		if len(args) > 0 {
			version = args[0]
		}

		latestVersion := "0.0.0"
		v, nil := semver.NewVersion(latestVersion)
		out, err = exec.Command("git", "tag").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		if len(out) > 0 {
			tags := strings.Split(strings.Trim(fmt.Sprintf("%s", out), "\n"), "\n")
			if len(tags) == 1 && tags[0] == "" {
				tags = []string{}
			}
			vs := make([]*semver.Version, len(tags))
			for i, r := range tags {
				localV, err := semver.NewVersion(r)
				if err != nil || !strings.Contains(r, ".") {
					vs[i] = v
				} else {
					vs[i] = localV
				}
			}
			sort.Sort(semver.Collection(vs))
			v = vs[len(vs)-1]
			latestVersion = v.Original()
		}

		if "patch" == version || "minor" == version || "major" == version {
			if "patch" == version {
				v := v.IncPatch()
				version = v.String()
			}
			if "minor" == version {
				v := v.IncMinor()
				version = v.String()
			}
			if "major" == version {
				v := v.IncMajor()
				version = v.String()
			}
			versionParts := strings.Split(version, ".")
			latestVersionParts := strings.Split(latestVersion, ".")
			for i, p := range latestVersionParts {
				if p != "0" {
					versionParts[i] = strings.TrimRight(p, "123456789") + versionParts[i]
				}
			}
			version = strings.Join(versionParts, ".")
		}

		out, err = exec.Command("git", "tag", version, "master").CombinedOutput()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", out))
		}
		if len(out) > 0 {
			fmt.Printf("%s\n", out)
		} else {
			fmt.Printf("Tag '%s' created on branch '%s'\n", version, "master")
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
