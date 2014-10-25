package main

import (
	"testing"
	"time"
)

const (
	DataPath = "testdata"
)

func BenchmarkProjectsFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectsFiles(DataPath)
	}
}

func BenchmarkProjects(b *testing.B) {
	start := time.Time{}
	end := time.Time{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Projects(DataPath, start, end)
	}
}
