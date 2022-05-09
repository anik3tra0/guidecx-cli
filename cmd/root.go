package cmd

import (
	"errors"
	"fmt"
	"github.com/anik3tra0/guidecx-cli/model"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var (
	// VERSION version of CLI
	VERSION      = "0.0.1"
	VersionLabel = "alpha"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "guidecx-cli",
	Short: "GuideCX Command Line Interface",
	Long: `GuideCX Web User Interface requires patience.
For developers riding in the fast lane, this CLI will get the job done without testing your patience.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gcx",
	Long:  `All software has versions. This is gcx's`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("GCX v%s%s -- HEAD", VERSION, VersionLabel)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.guidecx-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func promptGetInput(pc model.PromptContent) string {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}
