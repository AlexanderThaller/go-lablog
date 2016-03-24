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
	"bytes"
	"io"
	"os"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/juju/errgo"

	"github.com/spf13/cobra"
)

var flagAddTimeStamp time.Time
var flagAddTimeStampRaw string
var flagAddAutoCommit bool

func init() {
	flagAddTimeStamp = time.Now()

	cmdAdd.PersistentFlags().StringVarP(&flagAddTimeStampRaw, "timestamp", "t",
		flagAddTimeStamp.String(), "The timestamp for which to record the note.")
	cmdAdd.PersistentFlags().BoolVarP(&flagAddAutoCommit, "commit", "c",
		true, "If true entries will be autocommited to the repository entries are in.")

	// note
	cmdAdd.AddCommand(cmdAddNote)

	// todo
	cmdAdd.AddCommand(cmdAddTodo)
	cmdAddTodo.AddCommand(cmdAddTodoActive)
	cmdAddTodo.AddCommand(cmdAddTodoInActive)

	RootCmd.AddCommand(cmdAdd)
}

var cmdAdd = &cobra.Command{
	Use:   "add [command]",
	Short: "Add a new entry to the log",
	Long:  `Add a new entry like a note or a todo to the log. You have to specify a project for which we want to record the log for.`,
	Run:   runCmdAdd,
}

func runCmdAdd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var cmdAddNote = &cobra.Command{
	Use:   "note",
	Short: "Add a new note to the log",
	Long:  `Add a new note to the log which can have a timestamp and an free form value for text.`,
	RunE:  runCmdAddNote,
}

func runCmdAddNote(cmd *cobra.Command, args []string) error {
	project, timestamp, value, err := helper.ArgsToEntryValues(args, flagAddTimeStamp, flagAddTimeStampRaw)
	if err != nil {
		return errgo.Notef(err, "can not convert args to entry usable values")
	}

	buffer := new(bytes.Buffer)

	if value != "-" {
		buffer.WriteString(value)
	}

	// If there is something piped in over stdin append it to the already set
	// value

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		io.Copy(buffer, os.Stdin)
	}

	note := data.Note{
		Value:     buffer.String(),
		TimeStamp: timestamp,
	}

	err = helper.RecordEntry(flagDataDir, project, note, flagAddAutoCommit)
	if err != nil {
		errgo.Notef(err, "can not record note to store")
	}

	return nil
}

var cmdAddTodo = &cobra.Command{
	Use:   "todo [command]",
	Short: "Add a new todo to the log",
	Long:  `Add a new todo to the log which can have a timestamp, a toggle state (if its active or not) and an free form value for text.`,
	Run:   runCmdAddTodo,
}

func runCmdAddTodo(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var cmdAddTodoActive = &cobra.Command{
	Use:   "active",
	Short: "Add a new todo to the log and mark it as active",
	Long:  `Add a new todo to the log which can have a timestamp, is marked as active and an free form value for text.`,
	RunE:  runCmdAddTodoActive,
}

func runCmdAddTodoActive(cmd *cobra.Command, args []string) error {
	project, todo, err := helper.ArgsToTodo(args, flagAddTimeStamp, flagAddTimeStampRaw)
	if err != nil {
		return errgo.Notef(err, "can not convert args to todo")
	}

	todo.Active = true

	err = helper.RecordEntry(flagDataDir, project, todo, flagAddAutoCommit)
	if err != nil {
		return errgo.Notef(err, "can not record todo to store")
	}

	return nil
}

var cmdAddTodoInActive = &cobra.Command{
	Use:   "inactive",
	Short: "Add a new todo to the log and mark it as inactive",
	Long:  `Add a new todo to the log which can have a timestamp, is marked as inactive and an free form value for text.`,
	RunE:  runCmdAddTodoInActive,
}

func runCmdAddTodoInActive(cmd *cobra.Command, args []string) error {
	project, todo, err := helper.ArgsToTodo(args, flagAddTimeStamp, flagAddTimeStampRaw)
	if err != nil {
		return errgo.Notef(err, "can not convert args to todo")
	}

	todo.Active = false

	err = helper.RecordEntry(flagDataDir, project, todo, flagAddAutoCommit)
	if err != nil {
		errgo.Notef(err, "can not record todo to store")
	}

	return nil
}
