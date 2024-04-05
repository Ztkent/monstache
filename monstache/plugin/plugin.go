package main

import (
	"os"

	"github.com/Ztkent/monstache/pkg/monstachemap"
	"github.com/sirupsen/logrus"
)

/*
 This is a Monstache Plugin, it will be run every time a new document is indexed (inserted/updated).
 You can implement any of the 4 supported plugin types:
	func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error)
	func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error)
	func Pipeline(ns string, changeStream bool) (stages []interface, err error)
	func Process(input*monstachemap.ProcessPluginInput) error

Their uses are explained here: https://rwynn.github.io/monstache-site/advanced/#golang
*/

var logger = logrus.New()

func init() {
	// Setup the logger, so it can be parsed by datadog
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
}

func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	doc := input.Document

	// Monstache has 2 sync options, direct and change stream.
	// Direct are read from the database, change stream are read from the oplog.
	if !input.IsDirect {
		logger.WithFields(logrus.Fields{
			"DocID":    doc["_id"],
			"IsDirect": input.IsDirect,
		}).Debug("Change Stream Document processing")
	}

	output = &monstachemap.MapperPluginOutput{Document: doc}
	return
}
