package cmd

import (
	"fmt"
	"github.com/anik3tra0/guidecx-cli/model"
	"github.com/anik3tra0/guidecx-cli/store"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	timeCmd = &cobra.Command{
		Use:   "time",
		Short: "CRUD on Time Management",
		Long: `Allows you to perform CRUD on Time Management on GuideCX Tasks.
FYI we only support addition of time to a task as that is what users really require.
`,
	}
)

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.AddCommand(newCmd)
}

func promptSelectProject(pc model.PromptContent) string {
	var projectNames []string
	projects := store.ProjectsStore.Projects
	for _, p := range projects {
		projectNames = append(projectNames, p.Name)
	}

	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.Label,
			Items: projectNames,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			projectNames = append(projectNames, result)
		}
	}

	if err == promptui.ErrAbort {
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		fmt.Println("- Good bye!")
		os.Exit(0)
	}

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func promptSelectUser(pc model.PromptContent) string {
	var userNames []string
	users := store.UsersStore.Users
	for _, u := range users {
		userNames = append(userNames, u.Email)
	}

	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.Label,
			Items: userNames,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			userNames = append(userNames, result)
		}
	}

	if err == promptui.ErrAbort {
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		fmt.Println("- Good bye!")
		os.Exit(0)
	}

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func promptSelectMilestone(pc model.PromptContent, milestones []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.Label,
			Items: milestones,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			milestones = append(milestones, result)
		}
	}

	if err == promptui.ErrAbort {
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		fmt.Println("- Good bye!")
		os.Exit(0)
	}

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func promptSelectTask(pc model.PromptContent, tasks []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.Label,
			Items:    tasks,
			AddLabel: "Create New Task",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			tasks = append(tasks, result)
		}
	}

	if err == promptui.ErrAbort {
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		fmt.Println("- Good bye!")
		os.Exit(0)
	}

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func promptYesNoInput() bool {
	prompt := promptui.Select{
		Label: "Billable Hours[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()

	if err == promptui.ErrAbort {
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		fmt.Println("- Good bye!")
		os.Exit(0)
	}

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result == "Yes"
}
