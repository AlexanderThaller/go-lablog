package main

import (
	"flag"
	"strconv"
	"strings"
	"time"

	"github.com/AlexanderThaller/dbfiles"
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/store"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

var flagDataDir string
var flagOutDir string

func init() {
	now.TimeFormats = append(now.TimeFormats, time.RFC3339Nano)
}

func main() {
	flag.StringVar(&flagDataDir, "datadir", ".lablog", "the path to the datadir to use as the source of files.")
	flag.StringVar(&flagOutDir, "outdir", "output", "the path to the output directroy in which the converted files will be saved.")

	flag.Parse()

	log.Debug("DataDir: ", flagDataDir)
	log.Debug("OutDir: ", flagOutDir)

	dbread := dbfiles.New()
	dbread.Structure = dbfiles.NewFlat()
	dbread.BaseDir = flagDataDir

	readKeys, err := dbread.Keys()
	if err != nil {
		log.Fatal(errgo.Notef(err, "can not get keys from datadir"))
	}

	store, err := store.NewFolderStore(flagOutDir)
	if err != nil {
		log.Fatal(errgo.Notef(err, "can not create new store"))
	}

	for _, key := range readKeys {
		log.Info("Converting key '", strings.Join(key, "."), "'")

		project, err := data.ParseProjectName(strings.Join(key, data.ProjectNameSepperator))
		if err != nil {
			log.Warning(errgo.Notef(err, "can not convert key to project name"))
			continue
		}

		log.Debug("Project: ", project)

		values, err := dbread.Get(key...)
		if err != nil {
			log.Warning(errgo.Notef(err, "can not get values for key '"+strings.Join(key, ".")+"'"))
			continue
		}

		log.Debug("Values: ", values)

		err = convertValues(store, project, values)
		if err != nil {
			log.Warning(errgo.Notef(err, "can no convert values for key '"+strings.Join(key, ".")+"'"))
			continue
		}

		log.Info("Converted key '", strings.Join(key, "."), "'")
	}
}

func convertValues(store store.Store, project data.ProjectName, values [][]string) error {
	for _, value := range values {
		if len(value) < 2 {
			return errgo.New("value length must be at least 2")
		}

		timestamp, err := now.Parse(value[0])
		if err != nil {
			return errgo.Notef(err, "can not parse timestamp of value")
		}

		switch value[1] {
		case "note":
			log.Debug("Saving note")
			note := data.Note{
				Value:     value[2],
				TimeStamp: timestamp,
			}

			err := store.AddEntry(project, note)
			if err != nil {
				return errgo.Notef(err, "can not save note to store")
			}

		case "todo":
			log.Debug("Saving todo")
			done, err := strconv.ParseBool(value[3])
			if err != nil {
				return errgo.Notef(err, "can not parse bool from value")
			}

			todo := data.Todo{
				Value:     value[2],
				TimeStamp: timestamp,
				Active:    !done,
			}

			err = store.AddEntry(project, todo)
			if err != nil {
				return errgo.Notef(err, "can not save note to store")
			}
		default:
			return errgo.New("do not know what to do with this type of value: " + value[1])
		}
	}

	return nil
}
