package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "git-op",
	Long: strings.TrimPrefix(strings.Replace(`
           _ __
    ____ _(_) /_   ____  ____
   / __ '/ / __/  / __ \/ __ \
  / /_/ / / /_   / /_/ / /_/ /
  \__, /_/\__/   \____/ .___/
 /____/              /_/

 Git branching tool`, "'", "`", -1), "\n"),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	rootCmd.Long += color.New(color.FgGreen).Sprintf(" v%s", version)

	rootCmd.SetOutput(color.Output)

	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgYellow).SprintFunc())
	cobra.AddTemplateFunc("StyleCommand", color.New(color.FgGreen).SprintFunc())
	cobra.AddTemplateFunc("StyleFlag", color.New(color.FgGreen).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
		`Flags:`, `{{StyleHeading "Flags:"}}`,
		`{{rpad .Name .NamePadding }}`, `{{rpad .Name .NamePadding | StyleCommand }}`,
		`{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}`, `{{.LocalFlags.FlagUsages | trimTrailingWhitespaces | StyleFlag}}`,
		`{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}`, `{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces | StyleFlag}}`,
	).Replace(usageTemplate)
	rootCmd.SetUsageTemplate(usageTemplate)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.git-op.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".git-op" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".git-op")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
