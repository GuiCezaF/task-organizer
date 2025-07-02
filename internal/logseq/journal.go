package logseq

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func (j *Journal) WriteJournal(issues []redmine.Issue) error {
	return nil
}
