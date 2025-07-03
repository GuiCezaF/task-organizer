package logseq

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/GuiCezaF/task-organizer/internal/redmine"
)

type Journal struct {
	Filename string
	Path     string
}

func (j *Journal) NewJournal() error {
	now := time.Now()
	j.Filename = now.Format("2006_01_02") + ".md"

	fullPath := filepath.Join(j.Path, j.Filename)

	_, statErr := os.Stat(fullPath)
	if statErr == nil {
		log.Println("O journal já existe:", fullPath)
		return nil
	}
	if !os.IsNotExist(statErr) {
		return fmt.Errorf("erro ao verificar o arquivo: %w", statErr)
	}

	if mkErr := os.MkdirAll(j.Path, 0755); mkErr != nil {
		return fmt.Errorf("erro ao criar diretório: %w", mkErr)
	}

	file, createErr := os.Create(fullPath)
	if createErr != nil {
		return fmt.Errorf("erro ao criar o arquivo: %w", createErr)
	}
	defer file.Close()

	return nil
}

func (j *Journal) WriteJournal(issues []redmine.Issue, url string) error {
	fileFullPath := filepath.Join(j.Path, j.Filename)

	file, err := os.OpenFile(fileFullPath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	priorityOrder := map[string]int{
		"4 - Alta":           1,
		"3 - Normal":         2,
		"2 - Baixa":          3,
		"0 - Não priorizado": 4,
	}

	sort.SliceStable(issues, func(i, j int) bool {
		return priorityOrder[issues[i].Priority.Name] < priorityOrder[issues[j].Priority.Name]
	})

	_, err = file.WriteString("\n- ## Tarefas do dia\n")
	if err != nil {
		return err
	}

	for _, task := range issues {
		formattedTask := fmt.Sprintf("\t- TODO -> %d/%s | %s | Prioridade %s | [LINK](%s/issues/%d) \n",
			task.ID, task.Subject, task.Project.Name, task.Priority.Name, url, task.ID)

		_, err := file.WriteString(formattedTask)
		if err != nil {
			return err
		}
	}

	return nil
}
