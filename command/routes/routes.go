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

package routes

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
	// Constant asset paths
	RailsSqlAssetPath := "data/sql/github/rails-routes.sql"
	NodeSqlAssetPath := "data/sql/github/nodejs-routes.sql"
	TomcatSqlAssetPath := "data/sql/github/tomcat-routes.sql"

	verboseOpt := c.GlobalBool("verbose")
	silentOpt := c.GlobalBool("silent")
	testOpt := c.GlobalBool("test")
	project := c.GlobalString("project")
	credentials := c.GlobalString("credentials")
	limitArg := c.String("limit")
	frameworks := c.String("frameworks")
	outputFile := c.String("output")

	// TODO: Create route based extraction functionality
	// log.Fatal("Feature to be implemented...")

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
		"Mode":       "Routes",
	}

	// If test mode is enabled, let's use the testing SQL queries
	// testing SQL queries use `sample_contents` or smaller datasets.
	// This is to avoid breaking the bank!
	if testOpt {
		RailsSqlAssetPath = "data/sql/github/rails-routes-test.sql"
		NodeSqlAssetPath = "data/sql/github/nodejs-routes-test.sql"
		TomcatSqlAssetPath = "data/sql/github/tomcat-routes-test.sql"
		log.WithFields(fields).Info("Running in test mode.")
	}

	// Store generated templates in a string slice, if no
	// source parameter is provided, the default is to
	// run all sources available in the below switch statement
	sqlTemplateStrings := make(map[string]string)
	frameworkList := strings.Split(frameworks, ",")
	for _, fw := range frameworkList {
		switch fw {
		case "rails":
			RailsSqlAsset, err := assets.Asset(RailsSqlAssetPath)
			if err != nil {
				log.Debug("SQL for Rails not found.")
			}
			RailsSQLString := string(RailsSqlAsset[:])
			RailsTemplate := fasttemplate.New(RailsSQLString, "{{", "}}")
			RailsCompiledSql := RailsTemplate.ExecuteString(map[string]interface{}{
				"limit": limitArg,
			})
			if verboseOpt {
				log.WithFields(fields).Infof("Compiled SQL Template: %s", RailsCompiledSql)
			}
			sqlTemplateStrings["rails"] = RailsCompiledSql
			log.WithFields(fields).Info("Generated SQL template for Rails routes.")
		case "nodejs":
			NodeSqlAsset, err := assets.Asset(NodeSqlAssetPath)
			if err != nil {
				log.Debug("SQL for NodeJS not found.")
			}
			NodeSQLString := string(NodeSqlAsset[:])
			NodeTemplate := fasttemplate.New(NodeSQLString, "{{", "}}")
			NodeCompiledSql := NodeTemplate.ExecuteString(map[string]interface{}{
				"limit": limitArg,
			})
			if verboseOpt {
				log.WithFields(fields).Infof("Compiled SQL Template: %s", NodeCompiledSql)
			}
			sqlTemplateStrings["nodejs"] = NodeCompiledSql
			log.WithFields(fields).Info("Generated SQL template for NodeJS routes.")
		case "tomcat":
			TomcatSqlAsset, err := assets.Asset(TomcatSqlAssetPath)
			if err != nil {
				log.Debug("SQL for Tomcat not found.")
			}
			TomcatSQLString := string(TomcatSqlAsset[:])
			TomcatTemplate := fasttemplate.New(TomcatSQLString, "{{", "}}")
			TomcatCompiledSql := TomcatTemplate.ExecuteString(map[string]interface{}{
				"limit": limitArg,
			})
			if verboseOpt {
				log.WithFields(fields).Infof("Compiled SQL Template: %s", TomcatCompiledSql)
			}
			sqlTemplateStrings["tomcat"] = TomcatCompiledSql
			log.WithFields(fields).Info("Generated SQL template for Tomcat routes.")
		}
	}

	// Initialise BigQuery Golang Client
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, project, option.WithCredentialsFile(credentials))
	if err != nil {
		log.WithFields(fields).Warning("Failed to create a BigQuery client. Please check your credentials.")
	}

	// Iterate over generated templates and obtain results
	for framework, compiledSqlValue := range sqlTemplateStrings {
		fields["Framework"] = framework
		log.WithFields(fields).Info("Executing BigQuery SQL... this could take some time.")
		rows, err := query(client, ctx, compiledSqlValue)
		
		if err != nil {
			fields["Error"] = err.Error()
			log.WithFields(fields).Fatal("Error executing BigQuery SQL.")
		}
		
		resultsErr := handleResults(os.Stdout, rows, outputFile, framework, silentOpt, verboseOpt)
		
		if resultsErr != nil {
			fields["Error"] = resultsErr.Error()
			log.WithFields(fields).Fatal("Error handling results.")
		}
	}


	return nil
}