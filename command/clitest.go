package command

import (
	"github.com/urfave/cli"
	"log"
	"fmt"
)

func CmdStatus(c *cli.Context) error {
	project := c.String("project")

	if project == "" {
		log.Fatalln("project is required")
	}

	fmt.Printf("Initialising project: %s", project)

	return nil
}
