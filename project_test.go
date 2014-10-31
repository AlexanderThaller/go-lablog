package main

import (
	"bytes"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func BenchmarkProjectsFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectsFiles(TestDataPath)
	}
}

func BenchmarkProjects(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Projects(TestDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectWasActive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectWasActive(Project, TestDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectHasNotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectHasNotes(Project, TestDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectHasTodos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectHasTodos(Project, TestDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectNotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectNotes(Project, TestDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectNotesFromReader(b *testing.B) {
	path := path.Join(TestDataPath, Project+".csv")
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
		ProjectTodos(Project, TestDataPath, StartTime, EndTime)
	}
}
