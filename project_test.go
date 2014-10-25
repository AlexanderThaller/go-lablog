package main

import (
	"bytes"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

const (
	Project  = "Test1"
	DataPath = "testdata"
)

var (
	StartTime = time.Time{}
	EndTime   = time.Time{}
)

func BenchmarkProjectsFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectsFiles(DataPath)
	}
}

func BenchmarkProjects(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Projects(DataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectWasActive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectWasActive(Project, DataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectHasNotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectHasNotes(Project, DataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectHasTodos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectHasTodos(Project, DataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectNotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectNotes(Project, DataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectNotesFromReader(b *testing.B) {
	path := path.Join(DataPath, Project+".csv")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	buffer := bytes.NewBuffer(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProjectNotesFromReader(Project, StartTime, EndTime, buffer)
	}
}

func BenchmarkProjectTodos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectTodos(Project, DataPath, StartTime, EndTime)
	}
}
