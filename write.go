package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"time"
)

func WriteProjectNote(path string, timestamp time.Time, project, note string) error {
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
		timestamp.Format(time.RFC3339Nano),
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
