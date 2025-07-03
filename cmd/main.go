package main

import (
	"log"
	"os"

	"github.com/GuiCezaF/task-organizer/internal/logseq"
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
	logseq_path := os.Getenv("LOGSEQ_PATH")

	nj := logseq.Journal{Path: logseq_path}
	err = nj.NewJournal()

	if err != nil {
		log.Fatalf("Error in new journal: %s\n", err)
	}

	tasks, err := redmine.GetTasks(url, api_key)
	if err != nil {
		log.Fatal("Erro ao buscar tarefas:", err)
		return
	}

	err = nj.WriteJournal(tasks.Issue, url)
	if err != nil {
		log.Fatalf("Erro ao escrever tarefas no journal: %s", err)
	}
}
