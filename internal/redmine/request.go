package redmine

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RedmineIssueResponse struct {
	Issue []Issue `json:"issues"`
}

type Issue struct {
	ID             int           `json:"id"`
	Project        NamedEntity   `json:"project"`
	Tracker        NamedEntity   `json:"tracker"`
	Status         NamedEntity   `json:"status"`
	Priority       NamedEntity   `json:"priority"`
	Author         NamedEntity   `json:"author"`
	AssignedTo     *NamedEntity  `json:"assigned_to,omitempty"`
	Parent         *Parent       `json:"parent,omitempty"`
	Subject        string        `json:"subject"`
	Description    string        `json:"description"`
	StartDate      string        `json:"start_date"`
	DueDate        *string       `json:"due_date"`
	DoneRatio      int           `json:"done_ratio"`
	IsPrivate      bool          `json:"is_private"`
	EstimatedHours float64       `json:"estimated_hours"`
	CustomFields   []CustomField `json:"custom_fields"`
	CreatedOn      time.Time     `json:"created_on"`
	UpdatedOn      time.Time     `json:"updated_on"`
	ClosedOn       *time.Time    `json:"closed_on"`
}

type NamedEntity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Parent struct {
	ID int `json:"id"`
}

type CustomField struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func GetTasks(redmine_url string, api_key string) (*RedmineIssueResponse, error) {
	req, err := http.NewRequest("GET", redmine_url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("X-Redmine-API-Key", api_key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var tasks RedmineIssueResponse
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return &tasks, nil
}
