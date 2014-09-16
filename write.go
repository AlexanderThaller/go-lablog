package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"time"
)

func WriteProjectNote(path, project, note string) error {
	err := os.MkdirAll(path, 0750)
	if err != nil {
		return err
	}

	filepath := filepath.Join(path, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	values := []string{
		time.Now().Format(time.RFC3339Nano),
		CommandNoteString,
		note,
	}

	err = writer.Write(values)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}
