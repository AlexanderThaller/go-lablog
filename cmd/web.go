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
	"github.com/AlexanderThaller/lablog/src/web"
	log "github.com/Sirupsen/logrus"
	"github.com/juju/errgo"
	"github.com/spf13/cobra"
)

var (
	flagWebBinding string
)

func init() {
	webCmd.PersistentFlags().StringVarP(&flagWebBinding, "binding", "b",
		":18080", "The address and port to bind the webserver to.")
	webCmd.PersistentFlags().BoolVarP(&flagAddAutoCommit, "commit", "c",
		true, "If true entries will be autocommited to the repository entries are in.")

	RootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run an http webserver and serve http rendered version of the entries",
	RunE:  runWeb,
}

func runWeb(cmd *cobra.Command, args []string) error {
	level, err := log.ParseLevel(flagLogLevel)
	if err != nil {
		return errgo.Notef(err, "can not parse loglevel from flag")
	}

	err = web.Listen(flagDataDir, flagWebBinding, level)
	if err != nil {
		return errgo.Notef(err, "can not start web listener")
	}

	return nil
}
