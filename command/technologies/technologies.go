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

package technologies

import (
	"fmt"
	"os"
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
	technologiesAssetPath := "data/sql/http-archive/technologies.sql"
	verboseOpt := c.GlobalBool("verbose")
	silentOpt := c.GlobalBool("silent")
	project := c.GlobalString("project")
	credentials := c.GlobalString("credentials")
	limitArg := c.String("limit")
	technology := c.String("technology")
	date := c.String("date")
	outputFile := c.String("output")

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
	fields := log.Fields{
		"Mode":   "Technologies",
		"Source": "HTTPArchive",
		"Limit":  limitArg,
	}

	technologiesSqlAsset, err := assets.Asset(technologiesAssetPath)
	if err != nil {
		// Asset was not found.
	}
	currentTime := time.Now()
	haDate := currentTime.Format("2006_01")
	firstHaDate := fmt.Sprintf("%s_01", haDate)
	if date != "" {
		firstHaDate = date
	}
	wordsTemplate := string(technologiesSqlAsset[:])
	template := fasttemplate.New(wordsTemplate, "{{", "}}")
	compiledSql := template.ExecuteString(map[string]interface{}{
		"limit":      limitArg,
		"date":       firstHaDate,
		"technology": technology,
	})

	if verboseOpt {
		log.WithFields(fields).Infof("Compiled SQL Template: %s", compiledSql)
	}

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, project, option.WithCredentialsFile(credentials))
	if err != nil {
		log.WithFields(fields).Warning("Failed to create a BigQuery client. Please check your credentials.")
	}

	log.WithFields(fields).Info("Executing BigQuery SQL... this could take some time.")
	rows, err := query(client, ctx, compiledSql)

	if err != nil {
		fields["Error"] = err.Error()
		log.WithFields(fields).Fatal("Error executing BigQuery SQL.")
	}

	resultsErr := handleResults(os.Stdout, rows, outputFile, silentOpt, verboseOpt)

	if resultsErr != nil {
		fields["Error"] = err.Error()
		log.WithFields(fields).Fatal("Error handling results.")
	}

	return nil
}
