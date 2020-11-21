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

package main

import (
	"fmt"
	"os"

	"github.com/assetnote/commonspeak2/command/apiroutes"
	"github.com/assetnote/commonspeak2/command/deletedfiles"
	"github.com/assetnote/commonspeak2/command/directories"
	"github.com/assetnote/commonspeak2/command/parameters"
	"github.com/assetnote/commonspeak2/command/routes"
	"github.com/assetnote/commonspeak2/command/subdomains"
	"github.com/assetnote/commonspeak2/command/technologies"
	"github.com/assetnote/commonspeak2/command/wordswithext"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "project, p",
		Value: "",
		Usage: "The Google Cloud Project to use for the queries.",
	},
	cli.StringFlag{
		Name:     "credentials, c",
		Usage:    "Refer to: https://cloud.google.com/bigquery/docs/reference/libraries#client-libraries-install-go",
		FilePath: "credentials.json",
	},
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "Enable verbose output.",
	},
	cli.BoolFlag{
		Name:  "silent, s",
		Usage: "If this is set to true, the results will be written to a file but not to STDOUT.",
	},
	cli.BoolFlag{
		Name:  "test, t",
		Usage: "If this is set to true, Commonspeak2 will execute queries against smaller, testing datasets.",
	},
}

var Commands = []cli.Command{
	{
		Name:   "ext-wordlist",
		Usage:  "Generate wordlists based on extensions provided by the user.",
		Action: wordswithext.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "extensions, e",
				Value: "",
				Usage: "Extensions to generate wordlists e.g. html,php,js",
			},
			cli.StringFlag{
				Name:  "limit, l",
				Value: "200000",
				Usage: "Limit the wordlist to a certain number of lines. e.g. 200000",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Data output location e.g. wordlist.txt",
			},
			cli.StringFlag{
				Name:  "sources",
				Value: "github,httparchive",
				Usage: "Comma delimited sources to pull data from [github,httparchive]",
			},
			cli.StringFlag{
				Name:  "date",
				Value: "",
				Usage: "Date to use for HTTPArchive",
			},
		},
	},
	{
		Name:   "directories",
		Usage:  "Generate directory wordlists using URI.js and BigQuery",
		Action: directories.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "limit, l",
				Value: "200000",
				Usage: "Limit the wordlist to a certain number of lines. e.g. 200000",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Data output location e.g. wordlist.txt",
			},
			cli.StringFlag{
				Name:  "sources",
				Value: "httparchive",
				Usage: "Comma delimited sources to pull data from [httparchive]",
			},
			cli.StringFlag{
				Name:  "date",
				Value: "",
				Usage: "Date to use for HTTPArchive",
			},
		},
	},
	{
		Name:   "apiroutes",
		Usage:  "Generate API routes wordlists using URI.js and BigQuery",
		Action: apiroutes.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "limit, l",
				Value: "200000",
				Usage: "Limit the wordlist to a certain number of lines. e.g. 200000",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Data output location e.g. wordlist.txt",
			},
			cli.StringFlag{
				Name:  "sources",
				Value: "httparchive",
				Usage: "Comma delimited sources to pull data from [httparchive]",
			},
			cli.StringFlag{
				Name:  "date",
				Value: "",
				Usage: "Date to use for HTTPArchive",
			},
		},
	},
	{
		Name:   "parameters",
		Usage:  "Generate GET parameters wordlists using URI.js and BigQuery",
		Action: parameters.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "limit, l",
				Value: "200000",
				Usage: "Limit the wordlist to a certain number of lines. e.g. 200000",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Data output location e.g. wordlist.txt",
			},
			cli.StringFlag{
				Name:  "sources",
				Value: "httparchive",
				Usage: "Comma delimited sources to pull data from [httparchive]",
			},
			cli.StringFlag{
				Name:  "date",
				Value: "",
				Usage: "Date to use for HTTPArchive",
			},
		},
	},
	{
		Name:   "subdomains",
		Usage:  "Generates a list of subdomains from all available BigQuery public datasets.",
		Action: subdomains.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "limit, l",
				Value: "20000000",
				Usage: "Limit the subdomain list to a certain number of lines. e.g. 20000000",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Data output location e.g. subdomains.txt",
			},
			cli.StringFlag{
				Name:  "sources",
				Value: "hackernews,httparchive",
				Usage: "Comma delimited sources to pull data from [hackernews,httparchive]",
			},
			cli.StringFlag{
				Name:  "date",
				Value: "",
				Usage: "Date to use for HTTPArchive",
			},
		},
	},
	{
		Name:   "technologies",
		Usage:  "Generates a list of assets that have a specific technology from the HTTPArchive dataset.",
		Action: technologies.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "limit, l",
				Value: "20000000",
				Usage: "Limit the asset list to a certain number of lines. e.g. 20000000",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Data output location e.g. technologies.txt",
			},
			cli.StringFlag{
				Name:  "technology",
				Value: "Adobe Experience Manaager",
				Usage: "Technology you wish to retrieve assets for.",
			},
			cli.StringFlag{
				Name:  "date",
				Value: "",
				Usage: "Date to use for HTTPArchive",
			},
		},
	},
	{
		Name:   "routes",
		Usage:  "Generate wordlists based on routes extracted from popular frameworks.",
		Action: routes.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "frameworks, f",
				Value: "rails,nodejs,tomcat",
				Usage: "Frameworks to generate wordlists for, currently limited to [rails,nodejs,tomcat]",
			},
			cli.StringFlag{
				Name:  "limit, l",
				Value: "200000",
				Usage: "Limit the wordlist to a certain number of lines. e.g. 200000",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Data output location e.g. routes.txt",
			},
		},
	},
	{
		Name:   "deleted-files",
		Usage:  "Generate a list of deleted files based on GitHub commit messages.",
		Action: deletedfiles.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "limit, l",
				Value: "200000",
				Usage: "Limit the wordlist to a certain number of lines. e.g. 200000",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Data output location e.g. deleted.txt",
			},
		},
	},
}

var before = func(context *cli.Context) error {
	debugEnabled := context.GlobalBool("debug")
	if debugEnabled {
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
		logrus.SetLevel(logrus.DebugLevel)
	}
	return nil
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
