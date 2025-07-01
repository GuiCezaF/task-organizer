package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GuiCezaF/task-organizer/internal/redmine"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	url := os.Getenv("REDMINE_URL")
	api_key := os.Getenv("API_KEY")

	tasks, err := redmine.GetTasks(url, api_key)
	if err != nil {
		fmt.Println("Erro ao buscar tarefas:", err)
		return
	}

	for _, issue := range tasks.Issue {
		fmt.Printf("ID: %d - Task: %s\n", issue.ID, issue.Subject)
	}
}
