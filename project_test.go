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
		ProjectsFiles(BenchDataPath)
	}
}

func BenchmarkProjects(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Projects(BenchDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectWasActive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectWasActive(Project, BenchDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectHasNotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectHasNotes(Project, BenchDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectHasTodos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectHasTodos(Project, BenchDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectNotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectNotes(Project, BenchDataPath, StartTime, EndTime)
	}
}

func BenchmarkProjectNotesFromReader(b *testing.B) {
	path := path.Join(BenchDataPath, Project+".csv")
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
		ProjectTodos(Project, BenchDataPath, StartTime, EndTime)
	}
}
