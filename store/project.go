package store

import (
	"github.com/anik3tra0/guidecx-cli/model"
	"time"
)

type GetProjectsResponse struct {
	Projects []model.Project        `json:"projects"`
	Metadata map[string]interface{} `json:"metadata"`
}

var ProjectsStore GetProjectsResponse

type GetTasksResponse struct {
	Tasks    []model.Task           `json:"tasks"`
	Metadata map[string]interface{} `json:"metadata"`
}

var TasksStore GetTasksResponse

var MilestonesStore []model.Milestone

type GetUsersResponse struct {
	Users    []model.User           `json:"users"`
	Metadata map[string]interface{} `json:"metadata"`
}

var UsersStore GetUsersResponse

type CreateTaskRequest struct {
	AssigneeId      string    `json:"assigneeId"`
	Description     string    `json:"description"`
	DueDate         time.Time `json:"dueDate"`
	EstimatedHours  int       `json:"estimatedHours"`
	EstimatedPoints int       `json:"estimatedPoints"`
	MilestoneId     string    `json:"milestoneId"`
	Name            string    `json:"name"`
	Priority        string    `json:"priority"`
	Responsibility  string    `json:"responsibility"`
	SortOrder       int       `json:"sortOrder"`
	StartDate       time.Time `json:"startDate"`
	Visibility      string    `json:"visibility"`
}

type CreateTaskResponse struct {
	Id              string    `json:"id"`
	AssigneeId      string    `json:"assigneeId"`
	Description     string    `json:"description"`
	DueDate         time.Time `json:"dueDate"`
	EstimatedHours  int       `json:"estimatedHours"`
	EstimatedPoints int       `json:"estimatedPoints"`
	MilestoneId     string    `json:"milestoneId"`
	Name            string    `json:"name"`
	Priority        string    `json:"priority"`
	ProjectId       string    `json:"projectId"`
	Responsibility  string    `json:"responsibility"`
	SortOrder       int       `json:"sortOrder"`
	StartDate       time.Time `json:"startDate"`
	Status          string    `json:"status"`
	Visibility      string    `json:"visibility"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type CreateTimeRecordRequest struct {
	DateOfWork  string  `json:"dateOfWork"`
	ProjectId   string  `json:"projectId"`
	TaskId      string  `json:"taskId"`
	UserId      string  `json:"userId"`
	Hours       float64 `json:"hours"`
	Billable    bool    `json:"billable"`
	Description string  `json:"description"`
}

type CreateTimeRecordResponse struct {
	Id          string    `json:"id"`
	Billable    bool      `json:"billable"`
	DateOfWork  time.Time `json:"dateOfWork"`
	Description string    `json:"description"`
	Hours       float64   `json:"hours"`
	ProjectId   string    `json:"projectId"`
	TaskId      string    `json:"taskId"`
	UserId      string    `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
