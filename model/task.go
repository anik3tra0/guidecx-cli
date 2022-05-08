package model

type Task struct {
	Id       string `json:"id"`
	Assignee struct {
		Id        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	} `json:"assignee"`
	Attachments []struct {
		Id           string `json:"id"`
		CreatedAt    string `json:"createdAt"`
		InternalOnly bool   `json:"internalOnly"`
		Name         string `json:"name"`
		Protected    bool   `json:"protected"`
		UpdatedAt    string `json:"updatedAt"`
		Url          string `json:"url"`
	} `json:"attachments"`
	CompletedAt    string `json:"completedAt"`
	CreatedAt      string `json:"createdAt"`
	DependencyId   string `json:"dependency_id"`
	DependencyType string `json:"dependency_type"`
	Description    string `json:"description"`
	DueOn          string `json:"dueOn"`
	Duration       int    `json:"duration"`
	IsEvent        bool   `json:"isEvent"`
	Type           string `json:"type"`
	ProjectId      string `json:"projectId"`
	MilestoneId    string `json:"milestoneId"`
	Name           string `json:"name"`
	Notes          []struct {
		Id           string `json:"id"`
		InternalOnly bool   `json:"internalOnly"`
		Text         string `json:"text"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		CreatedBy    struct {
			Id        string `json:"id"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
		} `json:"createdBy"`
	} `json:"notes"`
	Priority       string      `json:"priority"`
	Responsibility string      `json:"responsibility"`
	StartOn        interface{} `json:"startOn"`
	Status         string      `json:"status"`
	Subtasks       []struct {
		Id          string `json:"id"`
		CompletedAt string `json:"completedAt"`
		CompletedBy struct {
		} `json:"completedBy"`
		CreatedAt string `json:"createdAt"`
		Done      bool   `json:"done"`
		Name      string `json:"name"`
		SortOrder int    `json:"sortOrder"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"subtasks"`
	UpdatedAt string `json:"updatedAt"`
}
