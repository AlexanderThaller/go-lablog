package main

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/AlexanderThaller/logger"
)

func GetProjects(datapath string) ([]string, error) {
	l := logger.New(Name, "GetProjects")

	dir, err := ioutil.ReadDir(datapath)
	if err != nil {
		return []string{}, err
	}

	var out []string
	for _, d := range dir {
		file := d.Name()
		l.Debug("DirName: ", file)

		if strings.HasPrefix(file, ".") {
			l.Trace(file, " has prefix '.'")
			continue
		}

		ext := filepath.Ext(file)
		name := file[0 : len(file)-len(ext)]

		out = append(out, name)
	}

	sort.Strings(out)
	return out, nil
}

func GetProjectRecords(datapath, project string) ([][]string, error) {
	path := filepath.Join(datapath, project+".csv")
	file, err := os.OpenFile(path, os.O_RDONLY, 0640)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, err
}

func GetRecords(datapath string) ([][]string, error) {
	projects, err := GetProjects(datapath)
	if err != nil {
		return nil, err
	}

	var out [][]string
	for _, project := range projects {
		records, err := GetProjectRecords(datapath, project)
		if err != nil {
			return nil, err
		}

		for _, record := range records {
			out = append(out, record)
		}
	}

	return out, nil
}

func GetTimeStamps(datapath string, startTime, endTime time.Time) ([]time.Time, error) {
	records, err := GetRecords(datapath)
	if err != nil {
		return nil, err
	}

	var out []time.Time
	for _, record := range records {
		timestamp := record[0]
		converted, err := time.Parse(time.RFC3339Nano, timestamp)
		if err != nil {
			return nil, err
		}

		if !startTime.Equal(time.Time{}) {
			if startTime.After(converted) {
				continue
			}
		}

		if !endTime.Equal(time.Time{}) {
			if endTime.Before(converted) {
				continue
			}
		}

		out = append(out, converted)
	}

	return out, nil
}

func GetDates(datapath string, startTime, endTime time.Time) ([]string, error) {
	timestamps, err := GetTimeStamps(datapath, startTime, endTime)
	if err != nil {
		return nil, err
	}

	filter := make(map[string]struct{})
	for _, timestamp := range timestamps {
		date := timestamp.Format("2006-01-02")
		filter[date] = struct{}{}
	}

	var out []string
	for date := range filter {
		out = append(out, date)
	}
	sort.Strings(out)

	return out, nil
}
