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
	"github.com/assetnote/commonspeak2/assets"
	"github.com/assetnote/commonspeak2/log"
	"github.com/icrowley/fake"
	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"golang.org/x/net/context"
	"io"
	"os"
	"fmt"
	"strings"
	"regexp"
)

func query(client *bigquery.Client, ctx context.Context, compiledSql string) (*bigquery.RowIterator, error) {
	query := client.Query(compiledSql)
	return query.Read(ctx)
}


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := assets.Asset(path)
  if err != nil {
	return nil, err
  }
  filterString := string(file[:])
  lines := strings.Split(filterString, "\n")
  return lines, nil
}



func cleanPathData(route string, framework string, numericalFilter []string, stringFilter []string) string {
	switch framework {
	case "rails":
		// Initial route clean up
		railsCleaner := strings.NewReplacer(
					      "'", "",
					      "\"", "",
					      "(", "",
					      ")", "",
					      "*", ":",
		   			    )
		route = railsCleaner.Replace(route)
		// Make numerical replacements (random number 5 digits)
		numericalReplacer := strings.NewReplacer(numericalFilter...)
		route = numericalReplacer.Replace(route)
		// Make string replacements (random english word)
		stringReplacer := strings.NewReplacer(stringFilter...)
		route = stringReplacer.Replace(route)
		// Generic replacements: integers are more likely valid than strings
		// as most frameworks will treat or convert integers to strings anyways...
		re := regexp.MustCompile("(\\:[a-zA-Z0-9_]*)")
	    route = re.ReplaceAllString(route, fake.Digits())
	}

	// Conform to starting with a "/"
    if !strings.HasPrefix(route ,"/") {
    	route = "/" + route
    }
	return route
}


func handleResults(w io.Writer, iter *bigquery.RowIterator, outputFile string, framework string, silent bool, verbose bool) error {
	fields := log.Fields{
		"Mode":       "Routes",
		"Framework":  framework,
		"Source":     "Github",
	}

	// obtain numerical parameter filters
	numericalParameters, err := readLines("data/filters/numerical-parameters.txt")
	if err != nil {
		log.WithFields(fields).Debugf("Numerical parameters not found: %s", err.Error())
	}
	
	// obtain string parameter filters
	stringParameters, stringErr := readLines("data/filters/string-parameters.txt")
	if stringErr != nil {
		log.WithFields(fields).Debugf("String parameters not found: %s", stringErr.Error())
	}
	
	numericalFilter := []string{}
	stringFilter := []string{}

	for _, parameter := range numericalParameters {
		fmt.Println(parameter)
		switch framework {
		case "rails":
			railsParam := fmt.Sprintf(":%s", parameter)
			railsNumerical := fake.Digits()
			numericalFilter = append(numericalFilter, railsParam, railsNumerical)
		case "nodejs":
			continue
		case "tomcat":
			continue
		}
	}

	for _, parameter := range stringParameters {
		fmt.Println(parameter)
		switch framework {
		case "rails":
			railsParam := fmt.Sprintf(":%s", parameter)
			railsString := fake.Word()
			stringFilter = append(stringFilter, railsParam, railsString)
		case "nodejs":
			continue
		case "tomcat":
			continue
		}
	}

	// fmt.Printf("%v",numericalFilter)
	// fmt.Printf("%v",stringFilter)

	file, err := os.Create(outputFile)
    if err != nil {
    	fields["Filename"] = outputFile
    	fields["Error"] = err.Error()
        log.WithFields(fields).Fatal("Cannot create output file")
    }
    defer file.Close()
    totalRows := 0
	for {
		var row Routes
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

		parsedRoute := cleanPathData(row.Route.String(), framework, numericalFilter, stringFilter)

		// Save to output file
		fmt.Fprintf(file, "%s\n", parsedRoute)
		// Print to console if verbose mode is on
		if verbose {
			fmt.Fprintf(w, "Route: %s Count: %s\n", parsedRoute, row.RouteCount.String())
		}
		totalRows++
	}
}