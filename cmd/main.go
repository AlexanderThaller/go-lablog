// Copyright © 2016 Alexander Thaller <alexander@thaller.ws>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"path"

	log "github.com/Sirupsen/logrus"

	"github.com/AlexanderThaller/lablog/src/helper"

	"github.com/juju/errgo"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var flagDataDir string
var flagLogLevel string

func init() {
	homepath, err := homedir.Dir()
	helper.ErrExit(errgo.Notef(err, "can not get homepath"))

	datadir := path.Join(homepath, ".lablog")

	// cmdMain
	RootCmd.PersistentFlags().StringVarP(&flagDataDir, "datadir", "d",
		datadir, "The path to the datadir for retreiving and storing the data.")
	RootCmd.PersistentFlags().StringVarP(&flagLogLevel, "loglevel", "l",
		"info", "The loglevel for which to run in. Default is warn. There are panic, fatal, error, warn info and debug as levels.")
}

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:               "lablog [command]",
	Short:             "lablog makes taking notes and todos easy",
	Long:              `lablog orders notes and todos into projects and subprojects without dictating a specific format.`,
	PersistentPreRunE: setLogLevel,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		log.Debug(errgo.Details(err))
	}
}

func setLogLevel(cmd *cobra.Command, args []string) error {
	level, err := log.ParseLevel(flagLogLevel)
	if err != nil {
		return errgo.Notef(err, "can not parse loglevel from flag")
	}

	log.SetLevel(level)

	return nil
}
