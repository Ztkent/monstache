package plugin

import (
	"os"

	"github.com/Ztkent/monstache/pkg/monstachemap"
	"github.com/sirupsen/logrus"
)

/*
 This is a Monstache Plugin, it will be run every time a new document is indexed (inserted/updated/deleted).
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

// We can modify the document before it is indexed.
// This is useful for adding new fields, or modifying existing fields to match the ElasticSearch schema.
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

// This sees every document that is processed by Monstache, before it is mapped.
// BulkProcessor: can be used to add requests process to any index.
// ElasticClient: can directly query ElasticSearch.

// We can target any index, and add any request type based on the message we are receiving.
// These requests will not be mapped, and are sent directly to ElasticSearch.
// Useful thread: https://github.com/rwynn/monstache/issues/181#issuecomment-468719179
func Process(input *monstachemap.ProcessPluginInput) error {
	// req := elastic.NewBulkIndexRequest()
	// req.Index("Seperate Index")
	// req.Id(input.Document["_id"])
	// req.Doc("New Data")
	// input.ElasticBulkProcessor.Add(req)
	return nil
}

// Before mapping a document from mongo, we can filter based on a documents content.
func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error) {
	// doc := input.Document
	// if doc["age"].(int) < 10 {
	// 	return true, nil
	// }
	return true, nil
}

// This allows you to run aggregation pipelines on the data before it is mapped.
// These are Mongo aggregation pipelines, not ElasticSearch.
func Pipeline(ns string, changeStream bool) (stages []interface{}, err error) {
	// Example: Pipeline that will only direct index documents with an 'age' greater than 10.
	/*
		if !changeStream {
			directReadsAggregation := map[string]interface{}{
				"$match": map[string]interface{}{
					"age": map[string]interface{}{
						"$gt": 10,
					},
				},
			}
			stages = append(stages, directReadsAggregation)
		}
		return stages, nil
	*/
	return nil, nil
}
