package model

import "time"

type Project struct {
	ActiveMilestones []string    `json:"activeMilestones"`
	ApiProjectId     interface{} `json:"apiProjectId"`
	CashValue        int         `json:"cashValue"`
	CompletedAt      interface{} `json:"completedAt"`
	CreatedAt        time.Time   `json:"createdAt"`
	Customer         struct {
		Domain string `json:"domain"`
		Name   string `json:"name"`
	} `json:"customer"`
	DoneAt                 interface{} `json:"doneAt"`
	EndOn                  string      `json:"endOn"`
	ExternalId             interface{} `json:"externalId"`
	Id                     string      `json:"id"`
	IntegrationReferenceId string      `json:"integrationReferenceId"`
	Name                   string      `json:"name"`
	ProjectManager         struct {
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		Id        string `json:"id"`
		LastName  string `json:"lastName"`
	} `json:"projectManager"`
	ProjectedEndDate  string        `json:"projectedEndDate"`
	ReferringObjectId interface{}   `json:"referringObjectId"`
	StartOn           string        `json:"startOn"`
	Status            string        `json:"status"`
	Tags              []interface{} `json:"tags"`
	TasksCount        int           `json:"tasksCount"`
	UpdatedAt         time.Time     `json:"updatedAt"`
}
