package model

type Milestone struct {
	Id          string      `json:"id"`
	CompletedAt string      `json:"completedAt"`
	CreatedAt   string      `json:"createdAt"`
	Description string      `json:"description"`
	Duration    int         `json:"duration"`
	Name        string      `json:"name"`
	SortOrder   interface{} `json:"sortOrder"`
	TasksCount  int         `json:"tasksCount"`
	UpdatedAt   string      `json:"updatedAt"`
}
