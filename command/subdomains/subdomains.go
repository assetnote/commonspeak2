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

package subdomains

import (
	"github.com/urfave/cli"
	"github.com/assetnote/commonspeak2/assets"
	"github.com/assetnote/commonspeak2/log"
	"golang.org/x/net/context"
	"github.com/valyala/fasttemplate"
	"google.golang.org/api/option"
	"cloud.google.com/go/bigquery"
	"os"
	"strings"
)

func CmdStatus(c *cli.Context) error {
	verboseOpt := c.GlobalBool("verbose")
	silentOpt := c.GlobalBool("silent")
	project := c.GlobalString("project")
	credentials := c.GlobalString("credentials")
	limitArg := c.String("limit")
	sources := c.String("sources")
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
		"Mode":       "Subdomains",
	}

	// Store generated templates in a string slice, if no
	// source parameter is provided, the default is to
	// run all sources available in the below switch statement
	SQLTemplateStrings := make(map[string]string)
	sourceList := strings.Split(sources, ",")
	for _, s := range sourceList {
		switch s {
		case "hackernews":
			HNSubdomainsSqlAsset, err := assets.Asset("data/sql/hackernews/subdomains.sql")
			if err != nil {
				log.Debug("SQL for HackerNews not found.")
			}
			HNSQLString := string(HNSubdomainsSqlAsset[:])
			HNTemplate := fasttemplate.New(HNSQLString, "{{", "}}")
			// TODO: Complete feature to extract data from the runs tables
			// httpArchiveTablePrefix := generateTablePrefixForHA()
			HNCompiledSql := HNTemplate.ExecuteString(map[string]interface{}{
				"limit": limitArg,
			})
			if verboseOpt {
				log.WithFields(fields).Infof("Compiled SQL Template: %s", HNCompiledSql)
			}
			SQLTemplateStrings["hackernews"] = HNCompiledSql
			log.WithFields(fields).Info("Generated SQL template for HackerNews.")
		case "httparchive":
			HASubdomainsSqlAsset, err := assets.Asset("data/sql/http-archive/subdomains.sql")
			if err != nil {
				log.Debug("SQL for HTTPArchive not found.")
			}
			HASQLString := string(HASubdomainsSqlAsset[:])
			HATemplate := fasttemplate.New(HASQLString, "{{", "}}")
			// TODO: Complete feature to extract data from the runs tables
			// httpArchiveTablePrefix := generateTablePrefixForHA()
			HACompiledSql := HATemplate.ExecuteString(map[string]interface{}{
				"limit": limitArg,
			})
			if verboseOpt {
				log.WithFields(fields).Infof("Compiled SQL Template: %s", HACompiledSql)
			}
			SQLTemplateStrings["httparchive"] = HACompiledSql
			log.WithFields(fields).Info("Generated SQL template for HTTPArchive.")
		}
	}

	// Initialise BigQuery Golang Client
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, project, option.WithCredentialsFile(credentials))
	if err != nil {
		log.WithFields(fields).Warning("Failed to create a BigQuery client. Please check your credentials.")
	}

	// Iterate over generated templates and obtain results
	for source, compiledSqlValue := range SQLTemplateStrings {
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