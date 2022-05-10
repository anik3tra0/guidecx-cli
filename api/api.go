package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/anik3tra0/guidecx-cli/model"
	data "github.com/anik3tra0/guidecx-cli/store"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	guideCXV1ApiURL = "https://api.guidecx.com/api/v1"
	guideCXV2ApiURL = "https://api.guidecx.com/api/v2"
	notStarted      = "not_started"
	workingOnIt     = "working_on_it"
	stuck           = "stuck"
	signOff         = "sign_off"
	done            = "done"
	notApplicable   = "not_applicable"
	notScheduled    = "not_scheduled"
	scheduled       = "scheduled"
)

var guideCXApiKey, _ = os.LookupEnv("GUIDECX_API_KEY")

// GetProjects fetches all projects from your GuideCX Account
func GetProjects(client *http.Client) ([]model.Project, error) {
	apiURL := fmt.Sprintf("%s/projects", guideCXV1ApiURL)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", guideCXApiKey))

	// Get Projects Response
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Convert to bytes
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var projectsRespData data.GetProjectsResponse
	err = json.Unmarshal(bts, &projectsRespData)
	if err != nil {
		return nil, err
	}

	return projectsRespData.Projects, nil
}

// GetTasksByProjectID fetches all tasks for a project from your GuideCX Account
func GetTasksByProjectID(client *http.Client, projectID string, milestoneID string) ([]model.Task, error) {
	apiURL := fmt.Sprintf("%s/projects/%s/tasks?milestoneId=%s", guideCXV1ApiURL, projectID, milestoneID)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", guideCXApiKey))

	// Get Tasks per Project
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Convert to bytes
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tasksRespData data.GetTasksResponse
	err = json.Unmarshal(bts, &tasksRespData)
	if err != nil {
		return nil, err
	}

	return tasksRespData.Tasks, nil
}

// GetMilestonesByProjectID fetches all milestones for a project from your GuideCX Account
func GetMilestonesByProjectID(client *http.Client, projectID string) ([]model.Milestone, error) {
	apiURL := fmt.Sprintf("%s/projects/%s/milestones", guideCXV1ApiURL, projectID)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", guideCXApiKey))

	// Get Milestones per Project
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Convert to bytes
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var milestonesRespData []model.Milestone
	err = json.Unmarshal(bts, &milestonesRespData)
	if err != nil {
		return nil, err
	}

	return milestonesRespData, nil
}

// GetUsers fetches all users from your GuideCX Account
func GetUsers(client *http.Client) ([]model.User, error) {
	apiURL := fmt.Sprintf("%s/users", guideCXV2ApiURL)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", guideCXApiKey))

	// Get Milestones per Project
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Convert to bytes
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var usersRespData data.GetUsersResponse
	err = json.Unmarshal(bts, &usersRespData)
	if err != nil {
		return nil, err
	}

	return usersRespData.Users, nil
}

// FindOrCreateProject fetches a particular project or creates a new project from your GuideCX Account
func FindOrCreateProject(client *http.Client, projectName string) ([]byte, error) {
	var newProject []byte

	store := data.ProjectsStore.Projects

	for _, p := range store {
		if p.Name == projectName {
			newProject, _ = json.Marshal(p)
			break
		}
	}

	// ToDo: If project is found then return project, else create new project
	if newProject != nil {
		return newProject, nil
	}

	return nil, fmt.Errorf("Project %s not found", projectName)
}

// CreateTaskByProjectID fetches a particular task or creates a new task for a project from your GuideCX Account
func CreateTaskByProjectID(client *http.Client, ctr *data.CreateTaskRequest, projectID string) (data.CreateTaskResponse, error) {
	var task data.CreateTaskResponse
	reqBody, _ := json.Marshal(ctr)
	apiURL := fmt.Sprintf("%s/projects/%s/tasks", guideCXV2ApiURL, projectID)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return task, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", guideCXApiKey))
	req.Header.Add("Content-Type", "application/json")

	// Get Tasks per Project
	resp, err := client.Do(req)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()

	// Convert to bytes
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return task, err
	}

	err = json.Unmarshal(bts, &task)
	if err != nil {
		return task, err
	}

	return task, nil
}

// CreateTimeRecord adds time to a task
func CreateTimeRecord(client *http.Client, ctr data.CreateTimeRecordRequest) (data.CreateTimeRecordResponse, error) {
	var timeRecordResponse data.CreateTimeRecordResponse
	reqBody, _ := json.Marshal(ctr)
	log.Println(string(reqBody))
	apiURL := fmt.Sprintf("%s/time-records", guideCXV2ApiURL)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return timeRecordResponse, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", guideCXApiKey))
	req.Header.Add("Content-Type", "application/json")

	// Create Time Record for a Task
	resp, err := client.Do(req)
	if err != nil {
		return timeRecordResponse, err
	}

	defer resp.Body.Close()

	// Convert to bytes
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return timeRecordResponse, err
	}

	err = json.Unmarshal(bts, &timeRecordResponse)
	if err != nil {
		return timeRecordResponse, err
	}

	return timeRecordResponse, nil
}
