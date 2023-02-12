/*
   Copyright 2018 Assetnote

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package pathswithbody

import (
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/assetnote/commonspeak2/log"
	"github.com/assetnote/commonspeak2/noisey"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func query(client *bigquery.Client, ctx context.Context, compiledSql string) (*bigquery.RowIterator, error) {
	query := client.Query(compiledSql)
	return query.Read(ctx)
}

func handleResults(w io.Writer, iter *bigquery.RowIterator, outputFile string, source string, silent bool, verbose bool) error {
	fields := log.Fields{
		"Mode":   "PathsWithBody",
		"Source": source,
	}
	file, err := os.Create(outputFile)
	if err != nil {
		fields["Filename"] = outputFile
		fields["Error"] = err.Error()
		log.WithFields(fields).Fatal("Cannot create output file")
	}
	defer file.Close()
	totalRows := 0
	for {
		switch source {
		case "github":
			var row GithubPaths
			err := iter.Next(&row)
			if err == iterator.Done {
				if !silent {
					log.WithFields(fields).Infof("Total rows extracted %d.", totalRows)
				}
				return nil
			}
			if err != nil {
				return err
			}
			// dont save if its an MD5,SHA1,SHA256,SHA512 hash or other noise
			if noisey.IsNotNoisey(row.Path.String()) {
				// Save to output file
				fmt.Fprintf(file, "%s\n", row.Path)
				// Print to console if verbose mode is on
				if verbose {
					fmt.Fprintf(w, "Path: %s Count: %s\n", row.Path, row.PathCount.String())
				}
			}
			totalRows++
		case "httparchive":
			var row HTTPArchivePaths
			err := iter.Next(&row)
			if err == iterator.Done {
				if !silent {
					log.WithFields(fields).Infof("Total rows extracted %d.", totalRows)
				}
				return nil
			}
			if err != nil {
				return err
			}
			// dont save if its an MD5,SHA1,SHA256,SHA512 hash or other noise
			if noisey.IsNotNoisey(row.Path.String()) {
				// Save to output file
				fmt.Fprintf(file, "%s\n", row.Path)
				// Print to console if verbose mode is on
				if verbose {
					fmt.Fprintf(w, "Path: %s Count: %s\n", row.Path, row.PathCount.String())
				}
				totalRows++
			}

		}

	}

}
