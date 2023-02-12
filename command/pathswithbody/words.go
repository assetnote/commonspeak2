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
	"os"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/assetnote/commonspeak2/assets"
	"github.com/assetnote/commonspeak2/log"
	"github.com/urfave/cli"
	"github.com/valyala/fasttemplate"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func CmdStatus(c *cli.Context) error {
	verboseOpt := c.GlobalBool("verbose")
	silentOpt := c.GlobalBool("silent")
	project := c.GlobalString("project")
	credentials := c.GlobalString("credentials")
	regexArg := c.String("regex")
	limitArg := c.String("limit")
	outputFile := c.String("output")
	sources := c.String("sources")
	date := c.String("date")
	sampleRate := c.Float64("sample-rate")

	if _, err := os.Stat(credentials); os.IsNotExist(err) {
		log.Fatalln(`credentials file not found, either have credentials.json saved in this folder, or provide the
						the credentials file path via the --credentials parameter.`)
	}

	if credentials == "" {
		log.Fatalln("credentials are required, provide it using the --credentials parameter.")
	}

	if project == "" {
		log.Fatalln("project is required, provide it using the --project parameter.")
	}

	if regexArg == "" {
		log.Fatalln("Regex string is required, provide it using the --regex parameter.")
	}
	if sampleRate > 100.0 || sampleRate < 0 {
		log.Fatalln("Sample percentage needs to be between 0.0 and 100.0")
	}

	fields := log.Fields{
		"Mode": "PathsWithBody",
	}
	sourceList := strings.Split(sources, ",")
	sqlTemplateStrings := make(map[string]string)
	for _, s := range sourceList {
		switch s {
		case "github":
			ghPathsWithBodySqlAsset, err := assets.Asset("data/sql/github/paths-with-body.sql")
			if err != nil {
				log.Fatal("Asset not found, ensure the asset was built with go-bindata")
			}
			ghPathsWithBodyTemplate := string(ghPathsWithBodySqlAsset[:])
			ghTemplate := fasttemplate.New(ghPathsWithBodyTemplate, "{{", "}}")
			ghCompiledSql := ghTemplate.ExecuteString(map[string]interface{}{
				"regex":       regexArg,
				"limit":       limitArg,
				"sample_rate": sampleRate,
			})

			fields := log.Fields{
				"Mode":       "PathsWithBody",
				"Source":     "Github",
				"Limit":      limitArg,
				"Regex":      regexArg,
				"SampleRate": sampleRate,
			}

			if verboseOpt {
				log.WithFields(fields).Infof("Compiled SQL Template: %s", ghCompiledSql)
			}
			sqlTemplateStrings["github"] = ghCompiledSql
			log.WithFields(fields).Info("Generated SQL template for Github.")
		case "httparchive":
			haPathsWithBodySqlAsset, err := assets.Asset("data/sql/http-archive/paths-with-body.sql")
			if err != nil {
				log.Fatal("Asset not found, ensure the asset was built with go-bindata")
			}
			haPathsWithBodyTemplate := string(haPathsWithBodySqlAsset[:])
			haTemplate := fasttemplate.New(haPathsWithBodyTemplate, "{{", "}}")
			currentTime := time.Now()
			haDate := currentTime.Format("2006_01")
			firstHaDate := fmt.Sprintf("%s_01", haDate)
			if date != "" {
				firstHaDate = date
			}
			haCompiledSql := haTemplate.ExecuteString(map[string]interface{}{
				"regex":       regexArg,
				"limit":       limitArg,
				"date":        firstHaDate,
				"sample_rate": fmt.Sprintf("%f", sampleRate),
			})

			fields := log.Fields{
				"Mode":       "PathsWithBody",
				"Source":     "HTTPArchive",
				"Limit":      limitArg,
				"Regex":      regexArg,
				"SampleRate": fmt.Sprintf("%f", sampleRate),
			}

			if verboseOpt {
				log.WithFields(fields).Infof("Compiled SQL Template: %s", haCompiledSql)
			}
			sqlTemplateStrings["httparchive"] = haCompiledSql
			log.WithFields(fields).Info("Generated SQL template for HTTP Archive.")
		}
	}
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, project, option.WithCredentialsFile(credentials))
	if err != nil {
		log.WithFields(fields).Warning("Failed to create a BigQuery client. Please check your credentials.")
	}

	// Iterate over generated templates and obtain results
	for source, compiledSqlValue := range sqlTemplateStrings {
		fields["Source"] = source
		log.WithFields(fields).Info("Executing BigQuery SQL... this could take some time.")
		rows, err := query(client, ctx, compiledSqlValue)

		if err != nil {
			fields["Error"] = err.Error()
			log.WithFields(fields).Fatal("Error executing BigQuery SQL.")
		}

		resultsErr := handleResults(os.Stdout, rows, outputFile, source, silentOpt, verboseOpt)

		if resultsErr != nil {
			fields["Error"] = resultsErr.Error()
			log.WithFields(fields).Fatal("Error handling results.")
		}
	}

	return nil
}
