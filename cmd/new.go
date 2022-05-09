package cmd

import (
	"fmt"
	"github.com/anik3tra0/guidecx-cli/api"
	"github.com/anik3tra0/guidecx-cli/model"
	"github.com/anik3tra0/guidecx-cli/store"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	newCmd = &cobra.Command{
		Use:   "new",
		Short: "To add new entity or sub entity",
		Long: `This allows you to add a new project, milestone, task or time to a project. 
FYI it only supports adding of time to a project as that is what users ideally want`,
		Run: func(cmd *cobra.Command, args []string) {
			initData()
			addTimeToTask()
		},
	}
	colorGreen       = "\033[32m"
	colorRed         = "\033[31m"
	guideCXAPIClient = http.DefaultClient
)

func initData() {
	guideCXAPIClient = http.DefaultClient
	projects, err := api.GetProjects(guideCXAPIClient)
	if err != nil {
		log.Println(colorRed, err)
		os.Exit(1)
	}

	store.ProjectsStore.Projects = append(store.ProjectsStore.Projects, projects...)

	users, err := api.GetUsers(guideCXAPIClient)

	if err != nil {
		log.Println(colorRed, err)
		os.Exit(1)
	}

	store.UsersStore.Users = append(store.UsersStore.Users, users...)
}

func addTimeToTask() {
	// Get User from Selection
	userPromptContent := model.PromptContent{
		ErrorMsg: "Please select a user",
		Label:    "Please select a user",
	}

	user := promptSelectUser(userPromptContent)
	// Get User ID
	userID := findUserID(user)
	log.Println(colorGreen, "User ID:", userID)

	// Get Project from Selection
	projectPromptContent := model.PromptContent{
		ErrorMsg: "Please select a project",
		Label:    "Please select a project",
	}

	project := promptSelectProject(projectPromptContent)

	// Get Project ID
	projectID := findProjectID(project)
	log.Println(colorGreen, "Project ID:", projectID)

	// Find all milestones for a projectID
	milestones, err := api.GetMilestonesByProjectID(guideCXAPIClient, projectID)
	if err != nil {
		log.Println(colorRed, err)
		os.Exit(1)
	}

	// Add Milestones to store
	store.MilestonesStore = append(store.MilestonesStore, milestones...)

	// Obtain Milestone Names
	var milestoneNames []string
	for _, m := range milestones {
		milestoneNames = append(milestoneNames, m.Name)
	}

	// Get Milestone from Selection
	milestonePromptContent := model.PromptContent{
		ErrorMsg: "Please select a Milestone",
		Label:    "Please select a Milestone",
	}

	milestone := promptSelectMilestone(milestonePromptContent, milestoneNames)
	milestoneID := findMilestoneID(milestone)

	log.Println(colorGreen, "MilestoneID", milestoneID)

	var taskNames []string

	// Find all tasks for a projectID
	tasks, err := api.GetTasksByProjectID(guideCXAPIClient, projectID, milestoneID)
	if err != nil {
		log.Println(colorRed, err)
		os.Exit(1)
	}

	// If there are no tasks then direct user to create taskName workflow
	if len(tasks) > 0 {
		// Add Tasks to store
		store.TasksStore.Tasks = append(store.TasksStore.Tasks, tasks...)

		// Obtain Task Names
		for _, t := range store.TasksStore.Tasks {
			taskNames = append(taskNames, t.Name)
		}
	}

	// Get taskName from Selection
	taskPromptContent := model.PromptContent{
		ErrorMsg: "Please select a taskName or add a new taskName",
		Label:    "Please select a taskName or add a new taskName",
	}

	taskName := promptSelectTask(taskPromptContent, taskNames)
	taskObj := fetchTaskByName(projectID, milestoneID, taskName)

	if len(strings.TrimSpace(taskObj.Id)) == 0 {
		ctr := store.CreateTaskRequest{
			Name:           taskName,
			AssigneeId:     userID,
			MilestoneId:    milestoneID,
			EstimatedHours: 1,
			Priority:       "MEDIUM",
			Responsibility: "INTERNAL",
			StartDate:      time.Now(),
			DueDate:        time.Now().AddDate(0, 0, 3),
			Visibility:     "VISIBLE",
		}

		taskData, err := api.CreateTaskByProjectID(guideCXAPIClient, &ctr, projectID)
		if err != nil {
			log.Println(colorRed, "Something went wrong. Please try again")
			addTimeToTask()
		}

		taskObj.Id = taskData.Id
		taskObj.Name = taskData.Name
	}

	log.Println(colorGreen, "TaskID", taskObj.Id)

	// Get time(in hours) from Input
	timePromptContent := model.PromptContent{
		ErrorMsg: "Please enter time in hours. 15 mins is 0.25, 30 mins is 0.50, 45 mins is 0.75",
		Label:    "Please enter time in hours. 15 mins is 0.25, 30 mins is 0.50, 45 mins is 0.75",
	}

	hours := promptGetInput(timePromptContent)
	taskTimeRecord, err := strconv.ParseFloat(hours, 64)
	if err != nil {
		log.Println(colorRed, "NaN. Enter a numeric value. Floats are allowed. Start this flow again")
	}

	// Get billable boolean from Input
	billable := promptYesNoInput()

	timeRecordRequest := store.CreateTimeRecordRequest{
		DateOfWork:  time.Now().Format("2006-01-02"),
		ProjectId:   projectID,
		TaskId:      taskObj.Id,
		UserId:      userID,
		Hours:       taskTimeRecord,
		Billable:    billable,
		Description: "",
	}

	timeRecordResponse, err := api.CreateTimeRecord(guideCXAPIClient, timeRecordRequest)
	if err != nil {
		log.Println(colorRed, err)
		addTimeToTask()
	}

	if len(timeRecordResponse.Id) > 0 {
		log.Println(colorGreen, fmt.Sprintf("Successfully added %2f to Task: %s for Milestone: %s under Project: %s", timeRecordResponse.Hours, taskName, milestone, project))
	} else {
		log.Println(colorRed, "Something went wrong, Please try this flow again")
	}

}

func findUserID(userName string) string {
	var uid string

	for _, u := range store.UsersStore.Users {
		if u.Email == userName {
			uid = u.Id
			break
		}
	}

	return uid
}

func findProjectID(projectName string) string {
	var pid string

	for _, p := range store.ProjectsStore.Projects {
		if p.Name == projectName {
			pid = p.Id
			break
		}
	}

	return pid
}

func findMilestoneID(milestoneName string) string {
	var mid string

	for _, m := range store.MilestonesStore {
		if m.Name == milestoneName {
			mid = m.Id
			break
		}
	}

	return mid
}

func fetchTaskByName(projectID string, milestoneID string, taskName string) model.Task {
	var task model.Task

	if len(store.TasksStore.Tasks) > 0 {
		for _, t := range store.TasksStore.Tasks {
			if t.Name == taskName && t.ProjectId == projectID && t.MilestoneId == milestoneID {
				task = t
				break
			}
		}
	}

	if len(strings.TrimSpace(task.Id)) == 0 {
		task.Id = ""
		task.Name = taskName
		task.MilestoneId = milestoneID
		task.ProjectId = projectID
	}

	return task
}
