package cmd

import (
	"log"
	"os"

	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

// selfUpdateCmd represents the self update command
var selfUpdateCmd = &cobra.Command{
	Use:     "selfupdate",
	Aliases: []string{"self-update", "su"},
	Short:   "Self update to github latest release",
	Run: func(cmd *cobra.Command, args []string) {
		latest, found, err := selfupdate.DetectLatest("garex/git-op")
		if err != nil {
			log.Fatalln("Error occurred while detecting version:", err)
			return
		}

		if !found {
			log.Fatalln("Latest version is not found")
			return
		}
		log.Println("Found latest version", latest.Version)

		exe, err := os.Executable()
		if err != nil {
			log.Fatalln("Could not locate executable path")
			return
		}

		err = selfupdate.UpdateTo(latest.AssetURL, exe)
		if err != nil {
			log.Fatalln("Error occurred while updating binary:", err)
			return
		}

		log.Println("Successfully updated to version", latest.Version)
	},
}

func init() {
	rootCmd.AddCommand(selfUpdateCmd)
}
